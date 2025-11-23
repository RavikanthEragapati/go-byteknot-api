--
-- DDL for Article Website Database Schema (MySQL)
-- Using InnoDB for transactional safety and Foreign Key support
--

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
CREATE TABLE article_views (
                               article_view_id BIGINT AUTO_INCREMENT PRIMARY KEY,
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