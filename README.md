At the final task of the VIX Full Stack Developer program, you are instructed to create an API based on the given case. For this case, you are required to build an API for uploading and deleting images. The API should support the following methods: POST, GET, PUT, and DELETE.

### API Specifications:

**User Endpoint:**

- **POST:** `/users/register`  
  Use the following attributes:  
  - `ID` (primary key, required)  
  - `Username` (required)  
  - `Email` (unique & required)  
  - `Password` (required & minimum length: 6)  
  - Relation with the `Photo` model (using cascade constraints)  
  - `Created At` (timestamp)  
  - `Updated At` (timestamp)  

- **GET:** `/users/login`  
  Use email and password (both required).  

- **PUT:** `/users/:userId`  
  Update user information.  

- **DELETE:** `/users/:userId`  
  Delete a user.

**Photos Endpoint:**

- **POST:** `/photos`  
  Attributes:  
  - `ID`  
  - `Title`  
  - `Caption`  
  - `PhotoUrl`  
  - `UserID`  
  - Relation with the `User` model.  

- **GET:** `/photos`  
  Retrieve all photos.

- **PUT:** `/photos/:photoId`  
  Update photo details.

- **DELETE:** `/:photoId`  
  Delete a specific photo.  

---

### Requirements:
- Use the Go JWT tool for authorization: [https://github.com/dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go).  
- Ensure that only the user who created a photo can delete or update it.

---

### Environment:
The document structure/environment for the GoLang project should resemble the following:  

- **`app`**: Contains the struct definition for the `user` model, used for data handling and authentication.  
- **`controllers`**: Includes database logic such as models and queries.  
- **`database`**: Contains database configuration and is used for establishing connections and performing migrations.  
- **`helpers`**: Includes utility functions used across the project, such as JWT handling, bcrypt, and header values.  
- **`middlewares`**: Contains functions for JWT authentication, used to protect the API.  
- **`models`**: Defines database models and relationships.  
- **`router`**: Contains routing configurations for accessing the API endpoints.  
- **`go.mod`**: Used for managing packages and dependencies (libraries).  

---

### Tools:  

You can use the following tools:  

- **Gin Gonic Framework**: [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)  
- **Gorm**: [https://gorm.io/index.html](https://gorm.io/index.html)  
- **JWT Go**: [https://github.com/dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)  
- **Go Validator**: [http://github.com/asaskevich/govalidator](http://github.com/asaskevich/govalidator)  

For the database, use **PostgreSQL**.  
