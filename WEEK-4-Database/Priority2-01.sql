CREATE TABLE User (
	UserID INT AUTO_INCREMENT PRIMARY KEY,
	Name VARCHAR(100) NOT NULL,
	Address VARCHAR(255),
	Email VARCHAR(100) NOT NULL UNIQUE,
	Description TEXT,
	CV_Document BLOB
);

CREATE TABLE JobVacancy (
	JobID INT AUTO_INCREMENT PRIMARY KEY,
	Title VARCHAR(100) NOT NULL,
	Description TEXT,
	JobFieldID INT,
	Status ENUM('Open', 'Closed') NOT NULL,
	RecruiterID INT,
	FOREIGN KEY (JobFieldID) REFERENCES JobField(JobFieldID),
	FOREIGN KEY (RecruiterID) REFERENCES Recruiter(RecruiterID)
);

CREATE TABLE Application (
	ApplicationID INT AUTO_INCREMENT PRIMARY KEY,
	UserID INT,
	JobID INT,
	CoverLetter TEXT,
	Status ENUM('Applied', 'Interview', 'Rejected', 'Accepted') NOT NULL,
	FOREIGN KEY (UserID) REFERENCES User(UserID),
	FOREIGN KEY (JobID) REFERENCES JobVacancy(JobID)
);

CREATE TABLE Recruiter (
	RecruiterID INT AUTO_INCREMENT PRIMARY KEY,
	Name VARCHAR(100) NOT NULL,
	Email VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE JobField (
	JobFieldID INT AUTO_INCREMENT PRIMARY KEY,
	FieldName VARCHAR(50) NOT NULL
);
