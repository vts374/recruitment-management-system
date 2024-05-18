Create a backend server for a Recruitment Management System.

Requirements:

1. Users can create their profile and upload Resumes (only PDF and DOCX formats allowed).
2. Uploaded resumes will be processed using a third-party API, and relevant information will be stored in the database.
3. Admin users can create job openings.
4. Admin users can view all uploaded resumes and extracted data of applicants.
5. Applicants can view job openings.
6. Applicants can apply to job openings.

APIs:
• POST /signup: Create a profile on the system (Name, Email, Password, UserType (Admin/Applicant), Profile Headline, Address).
• POST /login: Authenticate users and return a JWT token upon successful validation.
• POST /uploadResume: Authenticated API for uploading resume files (only PDF or DOCX) of the applicant. Only Applicant type users can access this API.
• POST /admin/job: Authenticated API for creating job openings. Only Admin type users can access this API.
• GET /admin/job/{job_id}: Authenticated API for fetching information regarding a job opening. Returns details about the job opening and a list of applicants. Only Admin type users can access this API.
• GET /admin/applicants: Authenticated API for fetching a list of all users in the system. Only Admin type users can access this API.
• GET /admin/applicant/{applicant_id}: Authenticated API for fetching extracted data of an applicant. Only Admin type users can access this API.
• GET /jobs: Authenticated API for fetching job openings. All users can access this API.
• GET /jobs/apply?job_id={job_id}: Authenticated API for applying to a particular job. Only Applicant users are allowed to apply for jobs.

Usage instructions: Send a request to the API with the resume file (PDF or DOCX), it will return a JSON with all the scrapped information as shown below:

{
"education": [
{
"name": "Wharton School of the University of Pennsylvania",
"url": "http://dbpedia.org/page/Wharton_School_of_the_University_of_Pennsylvania"
},
{
"name": "Penn's College of Arts and Sciences",
"url": "http://dbpedia.org/page/University_of_Pennsylvania_School_of_Arts_and_Sciences"
}
],
"email": "elonmusk@teslamotors.com",
"experience": [
{
"dates": [
"2006-06"
],
"name": "Solar City"
},
{
"dates": [
"06-2002"
],
"name": "SpaceX",
"url": "http://spacex.com"
},
{
"dates": [
"03-1999",
"10-2002"
],
"name": "X.com and Paypal",
"url": "http://paypal.com"
}
],
"name": "Elon Musk",
"phone": "650 68100",
"skills": [
"Entrepreneurship",
"Innovation",
"Mars",
"Space",
"Electric Cars",
"Physics",
"Maths",
"Calculus",
"Distrupting Technologies"
]
}

//Deployment commands are
//install the dependencies
go mod init
go mod tidy
go run main.go
Use Postman or curl to test the endpoints.
