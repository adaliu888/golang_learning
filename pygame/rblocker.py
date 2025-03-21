import sys
import pygame
import random

# 初始化 Pygame
pygame.init()

# 设置屏幕大小
SCREEN_WIDTH = 400
SCREEN_HEIGHT = 500
screen = pygame.display.set_mode((SCREEN_WIDTH, SCREEN_HEIGHT))
pygame.display.set_caption("俄罗斯方块")

# 定义颜色
BLACK = (0, 0, 0)
WHITE = (255, 255, 255)
GREEN = (0, 255, 0)
RED = (255, 0, 0)

# 方块大小
BLOCK_SIZE = 20

# 游戏区域大小
GRID_WIDTH = SCREEN_WIDTH // BLOCK_SIZE
GRID_HEIGHT = SCREEN_HEIGHT // BLOCK_SIZE

# 方块形状
SHAPES = [
    [[1, 1, 1, 1]],  # I
    [[1, 1], [1, 1]],  # O
    [[1, 0], [1, 0], [1, 1]],  # L
    [[0, 1], [0, 1], [1, 1]],  # J
    [[0, 1, 0], [1, 1, 1]],  # T
    [[0, 1, 1], [1, 1, 0]],  # S
    [[1, 1, 0], [0, 1, 1]]  # Z
]

# 方块颜色
COLORS = [
    (0, 255, 255),  # I
    (255, 255, 0),  # O
    (255, 165, 0),  # L
    (0, 0, 255),  # J
    (128, 0, 128),  # T
    (0, 255, 0),  # S
    (255, 0, 0)  # Z
]

# 定义方块类
class Tetromino:
    def __init__(self, shape, color):
        self.shape = shape
        self.color = color
        self.x = GRID_WIDTH // 2 - len(shape[0]) // 2
        self.y = 0

    def rotate(self):
        self.shape = list(zip(*self.shape[::-1]))

    def move(self, dx, dy):
        self.x += dx
        self.y += dy

# 检查方块是否可以放置
def check_position(board, shape, offset):
    off_x, off_y = offset
    for y, row in enumerate(shape):
        for x, cell in enumerate(row):
            try:
                if cell and board[y + off_y][x + off_x]:
                    return False
            except IndexError:
                return False
    return True

# 锁定方块到游戏区域
def lock_to_board(board, tetromino):
    for y, row in enumerate(tetromino.shape):
        for x, cell in enumerate(row):
            if cell:
                board[tetromino.y + y][tetromino.x + x] = tetromino.color

# 清除行
def clear_lines(board):
    lines_cleared = 0
    for i, row in enumerate(board[:-1]):
        if 0 not in row:
            del board[i]
            board.insert(0, [0] * GRID_WIDTH)
            lines_cleared += 1
    return lines_cleared

# 游戏主函数
def main():
    board = [[0] * GRID_WIDTH for _ in range(GRID_HEIGHT)]
    current_tetromino = Tetromino(random.choice(SHAPES), random.choice(COLORS))
    clock = pygame.time.Clock()
    running = True
    score = 0

    while running:
        screen.fill(BLACK)
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                running = False
            if event.type == pygame.KEYDOWN:
                if event.key == pygame.K_LEFT:
                    if check_position(board, current_tetromino.shape, (current_tetromino.x - 1, current_tetromino.y)):
                        current_tetromino.move(-1, 0)
                if event.key == pygame.K_RIGHT:
                    if check_position(board, current_tetromino.shape, (current_tetromino.x + 1, current_tetromino.y)):
                        current_tetromino.move(1, 0)
                if event.key == pygame.K_DOWN:
                    if check_position(board, current_tetromino.shape, (current_tetromino.x, current_tetromino.y + 1)):
                        current_tetromino.move(0, 1)
                if event.key == pygame.K_UP:
                    current_tetromino.rotate()
                    if not check_position(board, current_tetromino.shape, (current_tetromino.x, current_tetromino.y)):
                        current_tetromino.rotate()
                        current_tetromino.rotate()
                        current_tetromino.rotate()

        # 方块下落
        if check_position(board, current_tetromino.shape, (current_tetromino.x, current_tetromino.y + 1)):
            current_tetromino.move(0, 1)
        else:
            lock_to_board(board, current_tetromino)
            score += clear_lines(board)
            current_tetromino = Tetromino(random.choice(SHAPES), random.choice(COLORS))
            if not check_position(board, current_tetromino.shape, (current_tetromino.x, current_tetromino.y)):
                running = False

        # 绘制游戏区域
        for y, row in enumerate(board):
            for x, cell in enumerate(row):
                if cell:
                    pygame.draw.rect(screen, cell, (x * BLOCK_SIZE, y * BLOCK_SIZE, BLOCK_SIZE, BLOCK_SIZE), 0)

        # 绘制当前方块
        for y, row in enumerate(current_tetromino.shape):
            for x, cell in enumerate(row):
                if cell:
                    pygame.draw.rect(screen, current_tetromino.color, ((current_tetromino.x + x) * BLOCK_SIZE, (current_tetromino.y + y) * BLOCK_SIZE, BLOCK_SIZE, BLOCK_SIZE), 0)

        # 显示分数
        font = pygame.font.Font(None, 36)
        score_text = font.render(f"Score: {score}", True, WHITE)
        screen.blit(score_text, (10, 10))

        pygame.display.flip()
        clock.tick(10)

    pygame.quit()
    sys.exit()

if __name__ == "__main__":
    main()
