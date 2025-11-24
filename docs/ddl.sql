--
-- DDL for Article Website Database Schema (MySQL)
-- Using InnoDB for transactional safety and Foreign Key support
--

use bk;

-- 1. `users` Table
CREATE TABLE users (
                       user_id BIGINT AUTO_INCREMENT PRIMARY KEY,
                       firebase_uid VARCHAR(128) UNIQUE,
                       username VARCHAR(50) NOT NULL,
                       email VARCHAR(100) UNIQUE NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB;

-- 2. `articles` Table
CREATE TABLE articles (
                          article_id BIGINT AUTO_INCREMENT PRIMARY KEY,
                          title VARCHAR(255) NOT NULL,
                          content TEXT,
                          author_id BIGINT NOT NULL,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    -- Foreign Key Constraint
                          FOREIGN KEY (author_id) REFERENCES users(user_id) ON DELETE RESTRICT
) ENGINE=InnoDB;

-- 3. `comments` Table
CREATE TABLE comments (
                          comment_id BIGINT AUTO_INCREMENT PRIMARY KEY,
                          article_id BIGINT NOT NULL,
                          user_id BIGINT NOT NULL,
                          content TEXT NOT NULL,
                          parent_comment_id BIGINT, -- For replies (NULL for top-level comments)
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- Foreign Key Constraints
                          FOREIGN KEY (article_id) REFERENCES articles(article_id) ON DELETE CASCADE,
                          FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
                          FOREIGN KEY (parent_comment_id) REFERENCES comments(comment_id) ON DELETE CASCADE
) ENGINE=InnoDB;

-- 4. `likes` Table (Junction Table)
CREATE TABLE likes (
                       like_id BIGINT AUTO_INCREMENT PRIMARY KEY,
                       article_id BIGINT NOT NULL,
                       user_id BIGINT NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- Ensure a user can only like an article once
                       UNIQUE KEY uc_article_user (article_id, user_id),
    -- Foreign Key Constraints
                       FOREIGN KEY (article_id) REFERENCES articles(article_id) ON DELETE CASCADE,
                       FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
) ENGINE=InnoDB;

-- 5. `article_views` Table (Visitor Tracking)
CREATE TABLE views (
                               view_id BIGINT AUTO_INCREMENT PRIMARY KEY,
                               article_id BIGINT NOT NULL,
                               user_id BIGINT, -- NULL if visitor is not logged in
                               session_id VARCHAR(128), -- Unique identifier for guests/sessions
                               viewed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- Only one view per logged-in user per article is counted.
                               UNIQUE KEY uc_view_article_user (article_id, user_id),
    -- Foreign Key Constraint
                               FOREIGN KEY (article_id) REFERENCES articles(article_id) ON DELETE CASCADE,
                               FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE SET NULL -- Set to NULL if user is deleted, keeping view count history
) ENGINE=InnoDB;
commit;

-- #######################################################################
-- ####################### Insert data into tables #######################
-- #######################################################################

insert into users (user_id, firebase_uid, username, email)
values (1001, 'abcde1001', 'e.ravikanth', 'e.ravikanth@live.com');

insert into articles (article_id, title, content, author_id)
values (1111,'The five levels of skill','# The five levels of skill\n\n* __Novice:__ Requires clear rules and instructions to perform tasks. They have little to no experience and struggle to understand context.\n\* __Advanced Beginner:__ Starts to recognize situational elements and can perform some routine tasks independently, but still relies on guidelines for complex issues.\n* __Competent:__ Develops a conceptual understanding and can troubleshoot problems on their own. They can work independently and apply knowledge to a wider range of tasks.\n* __Proficient:__ Has a deeper understanding and can handle complex challenges. They anticipate issues, use judgment, and can adapt to new situations effectively.\n* __Expert:__ Possesses an intuitive grasp of the subject matter, going beyond rules and procedures. They can innovate, solve novel problems, and often mentor others.\n ',1001);

insert  into comments (comment_id, article_id, user_id, content)
values  (99,1111,1001,'great article thank you' );

insert  into comments (comment_id, article_id, user_id, content, parent_comment_id)
values  (88,1111,1001,'thanks for the message :)', 99 );

insert into likes (like_id, article_id, user_id)
values (55, 1111,1001);

insert into views (view_id, article_id, user_id, session_id)
values (44,1111,1001,'asfdsf');

commit ;

-- #######################################################################
-- ####################### Select data into tables #######################
-- #######################################################################


select * from users;
select article_id, title, content, author_id, created_at, updated_at  from articles;
select * from comments;
select * from likes;
select * from views;
commit;

-- #####################################################################
-- ####################### DROP data into tables #######################
-- #####################################################################

drop table views;
drop table likes;
drop table comments;
drop table articles;
drop table users;
commit;