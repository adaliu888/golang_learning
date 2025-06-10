package routes

import (
	"encoding/json"
	"golang_learning/user_auth/backend/database"
	"golang_learning/user_auth/backend/models"
	"golang_learning/user_auth/backend/utils"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes 注册用户相关路由
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/register", registerHandler).Methods("POST")
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/doctors", registerDoctorHandler).Methods("POST")
	r.HandleFunc("/protected", protectedHandler).Methods("GET")

}

// registerHandler 处理用户注册
func registerHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 输入验证
	if user.Username == "" || user.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// 哈希密码
	if err := user.HashPassword(); err != nil {
		http.Error(w, "Could not hash password", http.StatusInternalServerError)
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 将注册事件发送到 Redis 消息队列
	// utils.RedisClient.LPush(utils.ctx, "user_events", "User registered: "+user.Username)

	utils.Logger.Infof("User registered: %s", user.Username) // 记录注册事件
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// loginHandler 处理用户登录
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 输入验证
	if user.Username == "" || user.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	var dbUser models.User
	if err := database.DB.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// 比较密码
	if err := dbUser.ComparePassword(user.Password); err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// 生成 JWT
	token, err := utils.GenerateJWT(dbUser.Username)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	utils.Logger.Infof("User logged in: %s", user.Username) // 记录登录事件
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// registerDoctorHandler 处理医生注册
func registerDoctorHandler(w http.ResponseWriter, r *http.Request) {
	var doctor models.Doctor
	if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 输入验证
	if doctor.Name == "" || doctor.Specialty == "" {
		http.Error(w, "Name and specialty are required", http.StatusBadRequest)
		return
	}

	// 假设我们已经有用户 ID，实际应用中需要从 JWT 中提取
	doctor.UserID = 1 // 这里需要根据实际情况设置 UserID

	if err := database.DB.Create(&doctor).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Logger.Infof("Doctor registered: %s", doctor.Name) // 记录医生注册事件
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(doctor)
}

// protectedHandler 受保护的路由
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	claims, err := utils.ValidateJWT(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	utils.Logger.Infof("Accessed protected route by user: %s", claims.Username) // 记录访问事件
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to the protected route!", "user": claims.Username})
}
