-- 插入 exam_paper 表数据
INSERT INTO exam_paper (id, teacher_id, add_time, name, duration, status, description) VALUES 
(1, 1, '2024-05-22 10:00:00', '数学期末考试', 120, 1, '期末数学考试'),
(2, 2, '2024-05-22 10:05:00', '物理期中考试', 90, 1, '期中物理考试'),
(3, 3, '2024-05-22 10:10:00', '化学月考', 60, 1, '月考化学考试');

-- 插入 exam_question 表数据
INSERT INTO exam_question (id, add_time, paper_id, question, options, score, correct_answer, analysis, type) VALUES 
(1, '2024-05-22 10:00:00', 1, '2+2等于多少？', '{"A": "3", "B": "4", "C": "5"}', 5, 'B', '基础数学题', 1),
(2, '2024-05-22 10:05:00', 1, '3+3等于多少？', '{"A": "6", "B": "7", "C": "8"}', 5, 'A', '基础数学题', 1),
(3, '2024-05-22 10:10:00', 2, '光速是多少？', '{"A": "3x10^8 m/s", "B": "3x10^6 m/s", "C": "3x10^5 m/s"}', 5, 'A', '物理题', 1);

-- 插入 exam_score 表数据
INSERT INTO exam_score (id, student_id, paper_id, grading_teacher_id, score, comments, add_time) VALUES 
(1, 1, 1, 1, 90, '表现很好', '2024-05-23 09:00:00'),
(2, 2, 2, 2, 85, '有待提高', '2024-05-23 09:05:00'),
(3, 3, 3, 3, 78, '还需努力', '2024-05-23 09:10:00');


-- 插入 user 表数据
INSERT INTO user (id, username, password, role, add_time, email, name, avatar, gender) VALUES 
(1, 'admin', 'password123', 'admin', '2024-05-22 10:00:00', 'admin@example.com', 'Admin', 'avatar_admin.png', '未知'),
(2, 'teacher1', 'password123', 'teacher', '2024-05-22 10:05:00', 'teacher1@example.com', '张三', 'avatar1.png', '男'),
(3, 'student1', 'password123', 'student', '2024-05-22 10:10:00', 'student1@example.com', '赵六', 'avatar2.png', '女'),
(4, 'teacher2', 'password123', 'teacher', '2024-05-22 10:15:00', 'teacher2@example.com', '李四', 'avatar3.png', '男'),
(5, 'teacher3', 'password123', 'teacher', '2024-05-22 10:20:00', 'teacher3@example.com', '王五', 'avatar4.png', '男'),
(6, 'student2', 'password123', 'student', '2024-05-22 10:25:00', 'student2@example.com', '孙七', 'avatar5.png', '女'),
(7, 'student3', 'password123', 'student', '2024-05-22 10:30:00', 'student3@example.com', '周八', 'avatar6.png', '女');

-- 插入 teacher 表数据
INSERT INTO teacher (id, user_id, add_time, employee_id, department, id_number) VALUES 
(1, 2, '2024-05-22 10:00:00', 'T001', '数学系', '123456789'),
(2, 4, '2024-05-22 10:05:00', 'T002', '物理系', '987654321'),
(3, 5, '2024-05-22 10:10:00', 'T003', '化学系', '192837465');

-- 插入 student 表数据
INSERT INTO student (id, user_id, add_time, student_id, phone, id_number, address) VALUES 
(1, 3, '2024-05-22 10:00:00', 'S001', '1234567890', '987654321', '北京市海淀区'),
(2, 6, '2024-05-22 10:05:00', 'S002', '0987654321', '123456789', '上海市浦东新区'),
(3, 7, '2024-05-22 10:10:00', 'S003', '1122334455', '112233445', '广州市天河区');
