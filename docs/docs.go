package docs

import (
	"github.com/swaggo/swag"
)

var SwaggerInfo = &swag.Spec{
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

func InitSwagger() {
	SwaggerInfo.Title = "API E-Absensi"
	SwaggerInfo.Version = "1.0"
	SwaggerInfo.Schemes = []string{"http", "https"}
	SwaggerInfo.BasePath = "/"
}

const docTemplate = `

openapi: 3.0.2
info:
  version: 1.0.0
  title: API Immersive Dashboard
  description: An API for managing teams.

paths:
####################################
#  AUTH
####################################
  /auth/login:
    post:
      tags:
        - Auth
      summary: Login
      description: Login akun 
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              example:
                email: "rischi@mail.com"
                password: "123456"
      responses:
        '200':
          description: Login success.
          content:
            application/json:
              schema:
                type: object
                example:
                  status: true
                  message: "Login success"
                  data:
                      token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
        '401' :
          $ref: '#/components/responses/401'
        '400':
          $ref: '#/components/responses/400'

  /auth/users:
    get:
      tags:
        - Auth
      summary: Claim Token
      description: Claim Token
      security:
        - JWTAuth: [] 
      responses:
        '200':
          description: Claim success.
          content:
            application/json:
              schema:
                type: object
                example:
                  status: true
                  message: "-"
                  data:
                      id : 2
                      role : "user"
                      exp : 1678488436
        '401' :
          $ref: '#/components/responses/401'

  /auth/change-password:
    post:
      tags:
        - Auth
      summary: Change Password
      description: Change Password
      security:
        - JWTAuth: []
      requestBody:
        description: Change Password
        required: true
        content:
          application/json:
            schema:
              type: object
              required:
                - old_password
                - new_password
                - confirm_password
              properties:
                old_password:
                  type: string
                  description: old password
                new_password:
                  type: string
                  description: new password 
                confirm_password:
                  type: string
                  description: confirm password 
      responses:
        '200':
          description: Change Password success.
          content:
            application/json:
              schema:
                type: object
                example:
                  status: true
                  message: "-"
                  data: null
        '400':
          $ref: '#/components/responses/400'
        '401' :
          $ref: '#/components/responses/401'
        '500' :
          $ref: '#/components/responses/500'

####################################
#  USERS
####################################

  /users:
    get:
      tags:
        - Users
      summary: get all users
      description: Melihat seluruh data user

      responses:
        '200':
          description: A list of users
          content:
            application/json:
              schema:
                type: object
                example:
                  status: true
                  message: "success get all users"
                  data:
                    - id: 1
                      full_name: "reza"
                      email: "reza@mail.com"
                      address: "pekanbaru"
                    - id: 2
                      full_name: "rischi"
                      email: "rischi@mail.com"
                      address: "pekanbaru"
        '500':
          $ref: '#/components/responses/500'

    post:
      tags:
        - Users
      summary: add user
      description: Menambah data User (register)
      security:
        - JWTAuth: []
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              example:
                full_name: "reza yuda"
                email: "reza@mail.com"
                team_id: 2
                password: "qwerty"
                phone_number: "081234"
                address: "pekanbaru"
                role: "admin"
      responses:
        '201':
          description: Register berhasil.
          content:
            application/json:
              schema:
                type: object
                example:
                  status: true
                  message: "register berhasil"
        '401' :
          $ref: '#/components/responses/401'
        '500' :
          $ref: '#/components/responses/500'

  /users/{id}:
    parameters:
      - in: path
        name: id
        required: true
        description: Id
        schema:
          type: integer
    get:
      tags:
        - Users
      summary: get users by id
      security:
        - JWTAuth: []
      description: Mencari User dengan id

      responses:
        '200':
          description: A list of users
          content:
            application/json:
              schema:
                type: object
                example:
                  status: true
                  message: "success get all users"
                  data:
                    - id: 1
                      team_id: 2
                      full_name: "reza"
                      email: "reza@mail.com"
                      role : "admin"
                      status : true
                      phone_number : 0812345432
                      address: "land of dawn"
                      "team": {
                          "id": 1,
                          "name": "Placement"
                      }
        '404':
          $ref: '#/components/responses/404'
        '401' :
          $ref: '#/components/responses/401'
        '500' :
          $ref: '#/components/responses/500'

    put:
      tags:
        - Users
      summary: update users
      description: Mengubah data User
      security:
        - JWTAuth: []
      parameters:
        - name: user_id
          description: "id user"
          required: true
          in: path
          schema:
            type: integer
            example: 1
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              example:
                full_name: "reza yuda"
                email: "reza@mail.com"
                team_id: 2
                phone_number: "081234"
                address: "pekanbaru"
                status: false
                role: "admin"
      responses:
        '200':
          description: A list of users
          content:
            application/json:
              schema:
                type: object
                example:
                  status: true
                  message: "success Update users"
                  data: 
                    - id: 1
                      team_id: 2
                      full_name: "reza"
                      email: "reza@mail.com"
                      role : "admin"
                      status : true
                      phone_number : 0812345432
                      address: "land of dawn"
                      "team": {
                          "id": 1,
                          "name": "Placement"
                      }

        '404':
          $ref: '#/components/responses/404'
        '401' :
          $ref: '#/components/responses/401'
        '500' :
          $ref: '#/components/responses/500'
    
    delete:
      tags:
        - Users
      summary: Delete Users
      description: menghapus akun user
      security:
        - JWTAuth: []
      parameters:
        - name: "user_id"
          description: "user_id"
          required: true
          in: path
          schema:
            type: integer
            example: 1
      responses:
        '204':
          $ref: '#/components/responses/204'
        '404':
          $ref: '#/components/responses/404'
        '401' :
          $ref: '#/components/responses/401'
        '500' :
          $ref: '#/components/responses/500'

####################################
#  MENTEE
####################################

  /mentees:
    get:
      tags:
        - Mentee
      summary: get all mentees
      description: Melihat seluruh data mentees
      security:
        - JWTAuth: []
      parameters:
        - name: class
          description: "kelas mentee"
          required: false
          in: query
          schema:
            type: string
            example: "BE15"
        - name: status
          description: "status user"
          required: false
          in: query
          schema:
            type: string
            example: "On Class"
        - name: category
          required: false
          in: query
          schema: 
            type: string
            enum:
              - informatics
              - non-informatics
        - name: key
          description: ""
          required: false
          in: query
          schema:
            type: string
            example: key
      responses:
        '200':
          description: A list of users
          content:
            application/json:
              schema:

                type: object
                example:
                  status: true
                  message: "success get all mentees"
                  data:
                    - id: 1
                      class: "BE15"
                      full_name: "budi sanjaya"
                      nick_name: "budi"
                      status: "On Class"
                      education_type: non-informatics
                      email: "budi@mail.com"
                      gender: "L"
                    - id: 2
                      class: "BE15"
                      full_name: "budi doremi"
                      nick_name: "dora"
                      status: "On Class"
                      education_type: non-informatics
                      email: "budi@mail.com"
                      gender: "L"
        '404':
          $ref: '#/components/responses/404'
        '500':
          $ref: '#/components/responses/500'

    post:
      tags:
        - Mentee
      description: menambah data mentee
      summary: add Mentee
      requestBody:
        required: true
        content:
          application/json: 
            schema:
              type: object
              example:
                full_name: "budi sanjaya"
                nick_name: "budi"
                email: "budi@mail.com"
                phone: "0857777"
                current_address: "jl sesama"
                home_address: "jl sesama"
                telegram: "@busan"
                gender: "L"
                education_type: "non-informatics"
                major: "matematika"
                graduate: "2022"
                institution: "SMA 7 Pemalang"
                emergency_name: "bapakku"
                emergency_phone: "08967"
                emergency_status: "orang tua"
                status: "interview"
                class_id: 1
      responses:
        '201':
          description: Register berhasil.
          content:
            application/json:
              schema:
                type: object
                example:
                  status: true
                  message: "success register new mentee"
                  data:
                    id: 1
                    full_name: "budi sanjaya"
                    nick_name: "budi"
                    email: "budi@mail.com"
                    phone: "0857777"
                    telegram: "@busan"
                    gender: "L"        
        '404': 
          $ref: '#/components/responses/404'
        '500':
          $ref: '#/components/responses/500'

  /mentees/{id}:
    get: 
      tags:
        - Mentee
      summary: ini summary
      description: ini desc
      security:
        - JWTAuth: []
      parameters:
        - name: mentee_id
          description: "id mentee"
          required: true
          in: path
          schema: 
            type: integer
            example: 1
      responses:
        '200':
          description: data of specific mentee
          content:
            application/json:
              schema: 
                type: object
                example:
                  status: true
                  message: "success get specific mentee"
                  data:
                    id: 1
                    class: "BE15"
                    full_name: "budi sanjaya"
                    nick_name: "budi"
                    email: "budi@mail.com"
                    phone: "0857777"
                    current_address: "jl sesama"
                    home_address: "jl sesama"
                    telegram: "@busan"
                    gender: "L"
                    education_type: "non-informatics"
                    major: "matematika"
                    graduate: "2022"
                    institution: "SMA 7 Pemalang"
                    emergency_name: "bapakku"
                    emergency_phone: "08967"
                    emergency_status: "orang tua"
                    feedback: 
                      id: 1
                      notes: "Ini bagian notes untuk mentee yang bersangkutan"
                      users: "Jerry"
                      status: "Join"
                      proof: "https://images.app.goo.gl/t5b981L6oUTU3fE18"
        '404':
          $ref: '#/components/responses/404'
        '500':
          $ref: '#/components/responses/500'

    put:
      tags:
        - Mentee
      summary: "x"
      description: "edit data mentee"
      security:
        - JWTAuth: []
      parameters:
        - name: mentee_id
          description: "id mentee"
          required: true
          in: path
          schema: 
            type: integer
            example: 1
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              example:
                id: 1
                full_name: "budi sanjaya"
                nick_name: "budi"
                email: "budi@mail.com"
                phone: "0857777"
                telegram: "@busan"
                gender: "L" 
      responses:
        '200':
          description: edit specific mentee
          content:
            application/json:
              schema: 
                type: object
                example:
                  status: "succes"
                  message: "success update mentee"
        '404':
          $ref: '#/components/responses/404'
        '500':
          $ref: '#/components/responses/500'

    delete:
      tags:
        - Mentee
      summary: "delete mentee"
      description: "delete spesific mentee"
      security:
        - JWTAuth: []
      parameters:
        - name: mentee_id
          description: "id mentee"
          required: true
          in: path
          schema: 
            type: integer
            example: 1
      responses:
        '200':
          description: A list of users
          content:
            application/json:
              schema:
                type: object
                example:
                  status: true
                  message: "success delete mentee"      
        '404':
          $ref: '#/components/responses/404'
        '500':
          $ref: '#/components/responses/500'

  /mentees/{id}/feedbacks:
    get: 
      tags:
        - Mentee
      summary: ini summary
      description: ini desc
      security:
        - JWTAuth: []
      parameters:  
        - name: mentee_id
          description: "id mentee"
          required: true
          in: path
          schema: 
            type: integer
            example: 1
      responses:
        '200':
          description: "Sucess get feedback"
          content:
            application/json:
              schema:
                type: object
                # $ref: '#/components/schemas/ResponseObj'
              example:
                message: "Sucess update profile"
                data: 
                  mentee_id: 1
                  nama: "Jerry"
                  feedback:
                    - id: 1
                      notes: "Ini bagian notes untuk mentee yang bersangkutan 1"
                      users: "Jerry"
                      status: "Join"
                      proof: "https://images.app.goo.gl/t5b981L6oUTU3fE18"
                    - id: 3
                      notes: "Ini bagian notes untuk mentee yang bersangkutan 3"
                      users: "Jerry"
                      status: "Join"
                      proof: "https://images.app.goo.gl/t5b981L6oUTU3fE18"
        '404':
          $ref: '#/components/responses/404'
        '500':
          $ref: '#/components/responses/500'

####################################
#  FEEDBACK
####################################

  /feedbacks:
    post: 
      tags:
        - Feedback
      description: menambah data mentee
      summary: add Mentee
      requestBody:
        required: true
        content:
          application/json: 
            schema:
              type: object
              example:  
                notes: "Ini bagian notes untuk mentee yang bersangkutan 1"
                proof: "https://images.app.goo.gl/t5b981L6oUTU3fE18"
                user_id: 1
                mentee_id: 1
                status_id: 1
      responses:
        '201':
          description: Register berhasil.
          content:
            application/json:
              schema:
                type: object
                example:
                  status: true
                  message: "success register new mentee"  
                  data:
                    id: 1
                    notes: "Ini bagian notes untuk mentee yang bersangkutan 1"
                    users: "Jerry"
                    status: "Join"
                    proof: "https://images.app.goo.gl/t5b981L6oUTU3fE18"
        '404':
          $ref: '#/components/responses/404'
        '500':
          $ref: '#/components/responses/500'

  /feedbacks/{id}:
    put:
      tags:
        - Feedback
      summary: "x"
      description: "edit data feedback"
      security:
        - JWTAuth: []
      parameters:
        - name: id
          description: "id feedback"
          required: true
          in: path
          schema: 
            type: integer
            example: 1
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              example:
                  id: 1
                  full_name: "budi sanjaya"
                  nick_name: "budi"
                  email: "budi@mail.com"
                  phone: "0857777"
                  telegram: "@busan"
                  gender: "L"                
      responses:
        '200':
          description: edit specific mentee
          content:
            application/json:
              schema: 
                type: object
                example:
                  status: "succes"
                  message: "success update mentee"
        '404':
          $ref: '#/components/responses/404'
        '500':
          $ref: '#/components/responses/500'

    delete:
      tags:
        - Feedback
      summary: "delete feedback"
      description: "delete spesific feedback"
      security:
        - JWTAuth: []
      parameters:
        - name: id
          description: "id feedback"
          required: true
          in: path
          schema: 
            type: integer
            example: 1
      responses:
        '200':
          description: "-"
          content:
            application/json:
              schema:
                type: object
                example:
                  status: true
                  message: "success delete feedback"
        '404':
          $ref: '#/components/responses/404'
        '500':
          $ref: '#/components/responses/500'

####################################
#  TEAM
####################################

  /teams:
    get:
      tags:
      - Team
      summary: List all teams
      security:
        - JWTAuth: []
      responses:
        '200':
          description: A list of all teams
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/Team'
              
              examples:
                teamData:
                  description: "Example data team"
                  value: {
                    status: true,
                    message: "all teams",
                    data : [
                      {
                        id: "1",
                        name: "Placement",
                      },
                      {
                        id: "2",
                        name: "Academic",
                      },
                    ]
                  }
        
        '401' :
          $ref: '#/components/responses/401'
    post:
      tags:
        - Team
      summary: Create a new team
      security:
        - JWTAuth: []
      requestBody:
        description: The team to create
        required: true
        content:
          application/json:
            schema:
              type: object
              required:
                - name
              properties:
                name:
                  type: string
                  description: The name of the team
      responses:
        '201':
          description: The newly created team
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Team'

              examples:
                teamData:
                  description: "Example data team"
                  value: {
                    status: true,
                    message: "Create data success",
                    data : 
                      {
                        id: "1",
                        name: "Placement",
                      }
                  }               
        '400':
          $ref: '#/components/responses/400'
        '401' :
          $ref: '#/components/responses/401'
        '500' :
          $ref: '#/components/responses/500'
  /teams/{id}:
    parameters:
      - in: path
        name: id
        required: true
        description: The ID of the team to retrieve or update
        schema:
          type: integer

    get:
      tags:
      - Team
      summary: Get a team by ID
      security:
        - JWTAuth: []
      responses:
        '200':
          description: The requested team
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Team'
              
              examples:
                teamData:
                  description: "Example data team"
                  value: {
                    status: true,
                    message: "-",
                    data : 
                      {
                        id: "1",
                        name: "Placement",
                      }
                  }
        '404':
          $ref: '#/components/responses/404'
        '401' :
          $ref: '#/components/responses/401'
        '500' :
          $ref: '#/components/responses/500'
    
    put:
      tags:
      - Team
      summary: Update a team by ID
      security:
        - JWTAuth: []
      requestBody:
        description: The updated team
        required: true
        content:
          application/json:
            schema:
              type: object
              properties:
                name:
                  type: string
                  description: The updated name of the team
      responses:
        '200':
          description: The updated team
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Team'
              
              examples:
                teamData:
                  description: "Example data team"
                  value: {
                    status: true,
                    message: "Update data success",
                    data : 
                      {
                        id: "1",
                        name: "Placement",
                      }
                  }
        '404':
          $ref: '#/components/responses/404'
        '401' :
          $ref: '#/components/responses/401'
        '500' :
          $ref: '#/components/responses/500'  
    
    delete:
      tags:
      - Team
      summary: Delete a team by ID
      security:
        - JWTAuth: []
      responses:
        '204':
          $ref: '#/components/responses/204'
        '404':
          $ref: '#/components/responses/404'
        '401' :
          $ref: '#/components/responses/401'
        '500' :
          $ref: '#/components/responses/500'

####################################
#  CLASS
####################################

  /classes:
    get:
      tags:
        - Class
      summary: List all Class
      security:
        - JWTAuth: []
      responses:
        '200':
          description: A list of all teams
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/Class'
              
              examples:
                teamData:
                  description: "Example data class"
                  value: {
                    status: true,
                    message: "-",
                    data : [
                      {
                        "id": 1,
                        "name": "BE 11",
                        "pic": "Fakhry",
                        "start_date": "2022-06-10",
                        "graduate_date": "2022-08-30"
                      },
                      {
                        "id": 2,
                        "name": "FE 8",
                        "pic": "Bagas",
                        "start_date": "2022-06-10",
                        "graduate_date": "2022-08-30"
                      }
                    ]
                  }
        '404':
          $ref: '#/components/responses/404'
        '401' :
          $ref: '#/components/responses/401'
        '500' :
          $ref: '#/components/responses/500'

    post:
      tags:
        - Class
      summary: Create a new class
      security:
        - JWTAuth: []
      requestBody:
        description: The class to create
        required: true
        content:
          application/json:
            schema:
              type: object
              required:
                - name
                - pic
                - start_date
                - graduate_date
              properties:
                name:
                  type: string
                  description: The name of the class
                pic:
                  type: integer
                  description: pic
                start_date:
                  type: string
                  format : date
                  description: The name of the class
                graduate_date:
                  type: string
                  format : date
                  description: The name of the class
      responses:
        '201':
          description: The newly created class
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Class'

              examples:
                classData:
                  description: "Example data class"
                  value: {
                    status: true,
                    message: "Create data success",
                    data : 
                      {
                        "id": 1,
                        "name": "BE 11",
                        "pic": "Fakhry",
                        "start_date": "2022-06-10",
                        "graduate_date": "2022-08-30"
                      },
                  }
        '500' :
          $ref: '#/components/responses/500'
  /classes/{id}:
    parameters:
      - in: path
        name: id
        required: true
        description: The ID of the team to retrieve or update
        schema:
          type: integer

    get:
      tags:
      - Class
      summary: Get a class by ID
      security:
        - JWTAuth: []
      responses:
        '200':
          description: The requested class
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Class'
              
              examples:
                teamData:
                  description: "Example data class"
                  value: {
                    status: true,
                    message: "-",
                    data : 
                      {
                        "id": 1,
                        "name": "BE 11",
                        "pic" : 1,
                        "start_date": "2022-06-10",
                        "graduate_date": "2022-08-30"
                      }
                  }
        '400':
          $ref: '#/components/responses/400'
        '404':
          $ref: '#/components/responses/404'
        '401' :
          $ref: '#/components/responses/401'
        '500' :
          $ref: '#/components/responses/500'

    put:
      tags:
      - Class
      summary: Update a class by ID
      security:
        - JWTAuth: []
      requestBody:
        description: The updated class
        required: true
        content:
          application/json:
            schema:
              type: object
              properties:
                name:
                  type: string
                  description: The name of the class
                pic:
                  type: integer
                  description: pic
                start_date:
                  type: string
                  format : date
                  description: The name of the class
                graduate_date:
                  type: string
                  format : date
                  description: The name of the class
      responses:
        '200':
          description: The updated class
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Class'
              
              examples:
                teamData:
                  description: "Example data class"
                  value: {
                    status: true,
                    message: "Update data success",
                    data : 
                      {
                        id: "1",
                        name: "Placement",
                      }
                  }
        '400':
          $ref: '#/components/responses/400'
        '404':
          $ref: '#/components/responses/404'
        '401' :
          $ref: '#/components/responses/401'
        '500' :
          $ref: '#/components/responses/500'

    delete:
      tags:
      - Class
      summary: Delete a team by ID
      security:
        - JWTAuth: []
      responses:
        '204':
          description: The team was successfully deleted
          content:
            application/json:
              examples:
                teamData:
                  description: "Delete team"
                  value: {
                    status: true,
                    message: "Delete data success"
                  }
        '400':
          $ref: '#/components/responses/400'
        '404':
          $ref: '#/components/responses/404'
        '401' :
          $ref: '#/components/responses/401'
        '500' :
          $ref: '#/components/responses/500'

####################################
#  STATUS
####################################                
  /status/:
    get:
      tags:
        - Status
      summary: get all status
      description: get all status
      security:
        - JWTAuth: []
      parameters:
        - name: status_id
          description: "status"
          required: false
          in: query
          schema:
            type: string
            example: "budi"
  
      responses:
        '200':
          description: A list of status
          content:
            application/json:
              schema:
                type: object
                example:
                  status: true
                  message: "success get all status"
                  data:
                    - id: 1
                      name: "reza"
                      description: "pass to unit 2"
                      
        '404':
          $ref: '#/components/responses/404'

####################################
#  COMPONENT
####################################

components:
  securitySchemes:
    JWTAuth:
      type: http
      scheme: bearer
      bearerFormat: JWT
      description: "use Token"

  schemas:
    User:
      type: object
      properties:
        id:
          type: integer
          format: int64
        name:
          type: string
        email:
          type: string
        password:
          type: string
        phone_number:
          type: string
        address:
          type: string
        role:
          type: string
          enum:
            - admin
            - user
        id_team:
          type: integer
          enum : 
                [
                {
                 id: "1",
                 name: "Placement",
                },
                {
                 id: "2",
                 name: "Academic",
                }
                ]
        status:
          type: boolean
        created_at:
          type: string
          format: date-time
 
    Team:
      type: object
      required:
        - name
      properties:
        id:
          type: integer
          format: int64
          description: The ID of the team
        name:
          type: string
          description: The name of the team
        created_at:
          type: string
          format: date-time
          description: The date and time the team was created
        updated_at:
          type: string
          format: date-time
          description: The date and time the team was last updated
        deleted_at:
          type: string
          format: date-time
          description: The date and time the team was deleted
    
    Class:
      type: object
      required:
        - name
        - pic
        - start_date
        - graduate_date
      properties:
        id:
          type: integer
          format: int64
          description: The ID of the team
        name:
          type: string
          description: The name of the team
        pic:
          type: integer
          description: The name of the team
        start_date:
          type: string
          format: date
          description: The name of the team
        graduate_date:
          type: string
          format: date
          description: The name of the team
        created_at:
          type: string
          format: date-time
          description: The date and time the team was created
        updated_at:
          type: string
          format: date-time
          description: The date and time the team was last updated
        deleted_at:
          type: string
          format: date-time
          description: The date and time the team was deleted

    Feedback:
      type: object
      required:
        - notes
        - proof
        - user_id
        - mentee_id
        - status_id
      properties:
        id:
          type: integer
          format: int64
          description: The ID of the feedback
        notes:
          type: string
          description: The name of the notes
        proof:
          type: string
          description: The name of the proof
        user_id:
          type: integer
          format: int64
          description: The ID of the users
        mentee_id:
          type: integer
          format: int64
          description: The ID of the mentee
        status_id:
          type: integer
          format: int64
          description: The ID of the status
        created_at:
          type: string
          format: date-time
          description: The date and time the team was created
        updated_at:
          type: string
          format: date-time
          description: The date and time the team was last updated
        deleted_at:
          type: string
          format: date-time
          description: The date and time the team was deleted

    Mentee:
      type: object
      required:
        - class_id
        - full_name
        - nick_name
        - email
        - phone
        - current_address
        - home_address
        - telegram
        - status_id
        - gender
        - education_type
        - major
        - graduate
        - institution
        - emergency_name
        - emergency_phone
        - emergency_status
      properties:
        id:
          type: integer
          format: int64
          description: The ID of the feedback
        class_id:
          type: integer
          format: int64
          description: The ID of the users
        full_name:
          type: string
          description: The name of the full name
        nick_name:
          type: string
          description: The name of the nick name
        email:
          type: string
          description: email
        phone:
          type: string
          description: phone
        current_address:
          type: string
          description: current_address
        home_address:
          type: string
          description: home_address
        telegram:
          type: string
          description: telegram
        status_id:
          type: integer
          format: int64
          description: The ID of the status
        gender:
          type: string
          description: gender
        education_type:
          type: string
          description: education_type
        major:
          type: string
          description: major
        graduate:
          type: string
          description: graduate
        institution:
          type: string
          description: institution
        emergency_name:
          type: string
          description: emergency_name
        emergency_phone:
          type: string
          description: emergency_phone
        emergency_status:
          type: string
          description: emergency_status
        created_at:
          type: string
          format: date-time
          description: The date and time the team was created
        updated_at:
          type: string
          format: date-time
          description: The date and time the team was last updated
        deleted_at:
          type: string
          format: date-time
          description: The date and time the team was deleted

    Status:
      type: object
      required:
        - name
        - parent_id
        - description
      properties:
        id:
          type: integer
          format: int64
          description: The ID of the Status
        name:
          type: string
          description: The name of the Status
        parent_id:
          type: integer
          format: int64
          description: The ID of the Status which have be parent
        description:
          type: string
          description: description
        created_at:
          type: string
          format: date-time
          description: The date and time the team was created
        updated_at:
          type: string
          format: date-time
          description: The date and time the team was last updated
        deleted_at:
          type: string
          format: date-time
          description: The date and time the team was deleted
    
  responses:
    '400':
      description: Bad Request
      content:
        application/json:
          schema:
            type: object
            properties:
              status:
                type: boolean
              message:
                type: string
          example:
            status: false
            message: "error in your request"

    '401':
      description: Unauthorized
      content:
        application/json:
          schema:
            type: object
            properties:
              status:
                type: boolean
              message:
                type: string
          example:
            status: false
            message: "request does not contain a valid token"

    '404':
      description: Not Found
      content:
        application/json:
          schema:
            type: object
            properties:
              status:
                type: boolean
              message:
                type: string
          example:
            status: false
            message: "data not found"

    '204':
      description: Unauthorized
      content:
        application/json:
          schema:
            type: object
            properties:
              status:
                type: boolean
              message:
                type: string
          example:
            status: true
            message: "Delete data success"
            data : []

    '500':
      description: Internal Server Error
      content:
        application/json:
          schema:
            type: object
            properties:
              status:
                type: boolean
              message:
                type: string
          example:
            status: false
            message: "internal server error"


`
