-- exam_paper表
CREATE TABLE exam_paper (
    id BIGINT(20) NOT NULL AUTO_INCREMENT, -- 主键，自增ID
    teacher_id BIGINT(20) NOT NULL, -- 教师ID，对应 teacher 表中的 id
    add_time TIMESTAMP NULL, -- 添加时间
    name VARCHAR(200) NULL, -- 试卷名称
    duration INT(11) NULL, -- 考试时长（分钟）
    status INT(11) NULL, -- 状态（例如：0表示未发布，1表示已发布）
    description VARCHAR(500) NULL, -- 试卷描述
    PRIMARY KEY (id) -- 设置 id 为主键
);

-- exam_question表
CREATE TABLE exam_question (
    id BIGINT(20) NOT NULL AUTO_INCREMENT, -- 主键，自增ID
    add_time TIMESTAMP NULL, -- 添加时间
    paper_id BIGINT(20) NULL, -- 试卷ID，对应 exam_paper 表中的 id
    question VARCHAR(200) NULL, -- 试题内容
    options LONGTEXT NULL, -- 选项，JSON格式存储
    score BIGINT(20) NULL, -- 试题分数
    correct_answer VARCHAR(200) NULL, -- 正确答案
    analysis LONGTEXT NULL, -- 试题解析
    type INT(11) NULL, -- 试题类型（例如：1表示单选题，2表示多选题）
    PRIMARY KEY (id) -- 设置 id 为主键
);

-- 创建 exam_score 表
CREATE TABLE exam_score (
    id BIGINT(20) NOT NULL AUTO_INCREMENT,
    student_id BIGINT(20) NOT NULL, -- 对应的 student 表中的 id
    paper_id BIGINT(20) NOT NULL, -- 对应的 exampaper 表中的 id
    grading_teacher_id BIGINT(20) NOT NULL, -- 阅卷教师，对应的 teacher 表中的 id
    score INT(11) NOT NULL, -- 成绩
    comments VARCHAR(500) NULL, -- 阅卷老师的评论
    add_time TIMESTAMP NULL, -- 添加时间
    PRIMARY KEY (id)
);

-- 创建 user 表，包含所有用户的通用信息
CREATE TABLE user (
    id BIGINT(20) NOT NULL AUTO_INCREMENT,
    username VARCHAR(100) NULL, -- 用户名
    password VARCHAR(100) NULL, -- 密码
    role VARCHAR(100) NULL, -- 角色 (admin, teacher, student)
    add_time TIMESTAMP NULL, -- 添加时间
    email VARCHAR(200) NULL, -- 电子邮件
    name VARCHAR(200) NULL, -- 姓名
    avatar VARCHAR(200) NULL, -- 头像
    gender VARCHAR(50) NULL, -- 性别
    PRIMARY KEY (id),
    UNIQUE (username)
);

-- 创建 teacher 表，包含教师的特定信息
CREATE TABLE teacher (
    id BIGINT(20) NOT NULL AUTO_INCREMENT,
    user_id BIGINT(20) NOT NULL, -- 对应的 user 表中的 id
    add_time TIMESTAMP NULL, -- 添加时间
    employee_id VARCHAR(200) NULL, -- 教职工编号
    department VARCHAR(200) NULL, -- 部门
    id_number VARCHAR(200) NULL, -- 身份证号
    PRIMARY KEY (id),
);

-- 创建 student 表，包含学生的特定信息
CREATE TABLE student (
    id BIGINT(20) NOT NULL AUTO_INCREMENT,
    user_id BIGINT(20) NOT NULL, -- 对应的 user 表中的 id
    add_time TIMESTAMP NULL, -- 添加时间
    student_id VARCHAR(200) NULL, -- 学号
    phone VARCHAR(200) NULL, -- 电话
    id_number VARCHAR(200) NULL, -- 身份证号
    address VARCHAR(500) NULL, -- 地址
    PRIMARY KEY (id),
);
