CREATE TABLE users (
    userid INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    description TEXT
);


CREATE TABLE work_experiences (
    experienceid INT PRIMARY KEY AUTO_INCREMENT,
    userid INT,
    company_name VARCHAR(255) NOT NULL,
    job_title VARCHAR(255) NOT NULL,
    start_date DATE,
    end_date DATE,
    description TEXT,
    FOREIGN KEY (userid) REFERENCES users(userid)
);

CREATE TABLE recruiters (
    recruiterid INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    company VARCHAR(255)
);

CREATE TABLE job_vacancies (
    jobid INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    job_field VARCHAR(255),
    status VARCHAR(50),
    recruiterid INT,
    FOREIGN KEY (recruiterid) REFERENCES recruiters(recruiterid)
);

CREATE TABLE job_applications (
    applicationid INT PRIMARY KEY AUTO_INCREMENT,
    userid INT,
    jobid INT,
    cv TEXT,
    cover_letter TEXT,
    status VARCHAR(50),
    FOREIGN KEY (userid) REFERENCES users(userid),
    FOREIGN KEY (jobid) REFERENCES job_vacancies(jobid)
);

CREATE TABLE user_job_apps (
    userid INT,
    applicationid INT,
    PRIMARY KEY (userid, applicationid),
    FOREIGN KEY (userid) REFERENCES users(userid),
    FOREIGN KEY (applicationid) REFERENCES job_applications(applicationid)
);

