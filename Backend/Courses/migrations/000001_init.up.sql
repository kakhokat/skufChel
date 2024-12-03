create table users
(
    id bigserial primary key,
    email text not null unique ,
    password text not null ,
    name text not null,
    birthday date,
    image bytea,
    isCreator boolean,
    checkInt int,
    isConfirmed boolean
);

create table courses
(
    courseId bigserial primary key,
    name text not null,
    description text,
    likes int,
    creatorId int
);

create table courses_users
(
    id bigserial primary key,
    userId int references users(id),
    courseId int references courses(courseId) on delete cascade,
    status text,
    percent int,
    isLiked boolean
);

create table lessons
(
    lessonId bigserial primary key,
    name text not null,
    description text,
    testId int,
    videoId int,
    likes int
);

create table course_lesson
(
    courseId int references courses(courseId) on delete cascade,
    lessonId int references lessons(lessonId) on delete cascade,
    status text
);

create table users_lessons
(
    lessonId int references lessons(lessonId) on delete cascade,
    userId int references users(id),
    isPassed boolean,
    isLiked boolean
);

create table test
(
    testId bigserial primary key,
    question text,
    answers text,
    correctAnswer text
);

create table video
(
    videoId bigserial primary key,
    url text
)