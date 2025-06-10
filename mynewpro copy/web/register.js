document.addEventListener('DOMContentLoaded', function() {
    var registrationForm = document.getElementById('registrationForm');

    registrationForm.addEventListener('submit', function(event) {
        event.preventDefault(); // 阻止表单默认提交行为

        // 收集表单数据
        var username = document.getElementById('username').value;
        var password = document.getElementById('password').value;
        var email = document.getElementById('email').value;

        // 简单的前端验证
        if(username.trim() === '' || password.trim() === '' || email.trim() === '') {
            alert('所有字段都是必填项');
            return;
        }

        // 创建FormData对象
        var formData = new FormData(registrationForm);

        // 使用Fetch API发送AJAX请求
        fetch('/register', {
            method: 'POST',
            body: formData
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            console.log(data); // 处理返回的数据
            if(data.success) {
                alert('注册成功');
                // 可能需要跳转到其他页面或刷新当前页面
            } else {
                alert('注册失败: ' + data.message);
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });
    });
});