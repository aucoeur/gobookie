# BEW 2.5 MakeUtility Project Proposal (make.sc/makeutility)
   
  >**Name:** Aucoeur Ngo   
  >**Term:** 2021-T3

## 1. What problem will your MakeUtility solve?   
- What are your goals? Who will use your application?   

> Slide Deck: [Google Slides](https://docs.google.com/presentation/d/1Arqa2vxMgHoo7kf9O6k-Y0Q7y_OzU2oRHfp_OftALx4/view?usp=sharing)

## 2. What existing packages could help you build your project?   
- How does each library help you accomplish your goal(s)?   

   | 📦 Package<br />Name | 🏆 Helps Me Accomplish  |
   |---:|---|
   [bild](https://github.com/anthonynsimon/bild) | sobel edge detection 
   [cobra](https://cobra.dev/) | cli commands
   [image](https://golang.org/pkg/image/) | pre-processing image for edge detection

## 3. Break down your project into small tasks.   
- What steps will you take to implement the project?  
- Put them in order below.  Use as many spaces as you need to describe each step:

   ### 🔢 Task Description 
   - [x] Connect to GCP Cloud Vision API 
   - [x] Refactor initial annotations query from Node to Go
   - [x] Save annotations to JSON
   - [ ] Write tests
   - [ ] Integrate image processing techniques
     - [x] Research edge detection algos
     - [x] Research image packages
   - [ ] Save to DB
   - [ ] Use external API for book searches (or Product Search?)
   - [ ] Set up GUI/frontend
   - [ ] Deploy with Google App Engine (?)
