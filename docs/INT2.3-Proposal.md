# Seniors Spring Intensive Deliverable Proposal 

**Dates:** 3/16-3/25

**Name:** Aucoeur Ngo  
**Project Name:** gobookie  
**Is your project New or Old?** Old  
**Is your project Solo or Team?** Solo  

## Description

**Write a paragraph summary of the current status of your project, what you hope to achieve during the intensive, how and why**
gobookie aims to help users quickly catalog, check prices for resale, donation, etc in one shot. Take a photo of your bookshelves and gobookie will grab title/author information using Google's Cloud Vision API and return additional metadata and check their current resale value.

## Objective 1: Group text by book spine using image processing techniques for edge detection

**Why do you want to meet this objective? How will it improve your project?** 
Right now text annotation returns all text found without delineating between each book.  By splitting text up, the goal is to be able to identify and return product search data individually per book

**How will you demonstrate completion of your objective?** 
Text within book spine's bounding box will be grouped together and returned as one object

## Objective 2: Write tests
**Why do you want to meet this objective? How will it improve your project?** 
This will establish expected behavior and improve code quality and pave the way for automated testing pipelines 

**How will you demonstrate completion of your objective?** 
test files in repo with comments

## Objective 3: Design a frontend
**Why do you want to meet this objective? How will it improve your project?** 
Users can perform operations with a GUI. Having a frontend would make the app easier to use for the users and set the stage for making it available to use on mobile devices (arguably the defacto, easiest way for this to be used.. take a picture with your phone/device and upload for processing)

**How will you demonstrate completion of your objective?** 
Wireframes, MVP UI deployed to ghphages or github commit history

## Stretch goals (optional):

**What stretch goals do you have for your project?**
- Integrate with 3rd party API for searching titles
- Frontend in React Native
- Setup user accounts/auth and CRUD operations
    - This will set the structure for future users to track and modify their own collections and perhaps share/compare with other users.
- Automated test pipeline

## Evaluation

**You must meet the following criteria in order to pass the intensive:**

[rubric]:https://docs.google.com/document/d/1IOQDmohLBEBT-hyr-2vgw1mbZUNsq3fHxVfH0oRmVt0/edit
- Students must get proposal approved before starting the project to pass
- SOLO  
    - Must score an average above a 3 on the [rubric]  
- TEAM  
    - Must score an average above 3 on the [rubric]  
    - Each individual completes 2 of the 3 objectives from their proposal
- Pitch your product

## Approval Checklist
- [x] ~~If I have a team project,~~ I wrote this proposal to represent my work and only my work
- [x] I have completed all the necessary parts of this proposal
- [x] I linked my proposal in the Spring Intensive Tracker

### Sign off

> **Student Name:**                
> Aucoeur Ngo / March 15, 2021  
> **Make School Advisor Name**  
> TBD
