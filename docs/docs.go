package docs

import (
	"github.com/swaggo/swag"
)

const docTemplate = `

openapi: 3.0.2
info:
  version: 1.0.0
  title: Team Management API
  description: An API for managing teams.

paths:
  /login:
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
        '201':
          description: Login berhasil.
          content:
            application/json:
              schema:
                type: object
                example:
                  status: "success"
                  message: "Login berhasil"
                  data:
                      id: 1
                      email: "rischi@mail.com"
                      token:  "afkdnnifnnj&h3"
                      role: "Admin"
        '404' :
          $ref: '#/components/responses/404'

  /users:
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
                id_team: 2
                password: "qwerty"
                phone_number: "081234"
                address: "pekanbaru"
                status: "active"
                role: "admin"
      responses:
        '201':
          description: Register berhasil.
          content:
            application/json:
              schema:
                type: object
                example:
                  status: "success"
                  message: "register berhasil"
        '400':
          description: Register gagal.
          content:
            application/json:
              schema:
                type: object
                example:
                  status: "failed"
                  message: "register gagal. semua field wajib diisi"
        '500':
          $ref: '#/components/responses/500'

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
                  status: "success"
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
        '404':
          $ref: '#/components/responses/404'

  /users/{id_user}:
    get:
      tags:
        - Users
      summary: get users by name
      description: Mencari User dengan nama
      parameters:
        - name: name
          description: "nama user"
          required: false
          in: query
          schema:
            type: string
            example: "budi"

      responses:
        '200':
          description: A list of users
          content:
            application/json:
              schema:
                type: object
                example:
                  status: "success"
                  message: "success get all users"
                  data:
                    - id: 1
                      full_name: "reza"
                      email: "reza@mail.com"
                      address: "pekanbaru"
                    
        '404':
          $ref: '#/components/responses/404'

    put:
      tags:
        - Users
      summary: update users
      description: Mengubah data User
      security:
        - JWTAuth: []
      parameters:
        - name: id_user
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
                id_team: 2
                password: "qwerty"
                phone_number: "081234"
                address: "pekanbaru"
                status: "active"
                role: "admin"
      responses:
        '200':
          description: A list of users
          content:
            application/json:
              schema:
                type: object
                example:
                  status: "success"
                  message: "success Update users"
                  data: 
                    full_name: "reza yuda"
                    email: "reza@mail.com"
                    team: "be"
                    password: "qwerty"
                    phone_number: "081234"
                    address: "pekanbaru"
                    status: "active"
                    role: "admin"

        '404':
          $ref: '#/components/responses/404'
    
    delete:
      tags:
        - Users
      summary: Delete Users
      description: menghapus akun user
      security:
        - JWTAuth: []
      parameters:
        - name: "id_user"
          description: "id_user"
          required: true
          in: path
          schema:
            type: integer
            example: 1
      responses:
        '201':
          description: Delete berhasil.
          content:
            application/json:
              schema:
                type: object
                example:
                  status: "success"
                  message: "Delete success"
                  data:
                    full_name: reza
                    email: reza@mail.com
                    id_team: 2
        '404':
          $ref: '#/components/responses/404'
        '500':
          $ref: '#/components/responses/500'

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
                  status: "success"
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
                  status: "success"
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

  /mentees/{id_mentee}:
    get: 
      tags:
        - Mentee
      summary: ini summary
      description: ini desc
      security:
        - JWTAuth: []
      parameters:
        - name: id_mentee
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
                  status: "success"
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
        - name: id_mentee
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
        - name: id_mentee
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
                  status: "success"
                  message: "success delete mentee"      
        '404':
          $ref: '#/components/responses/404'
        '500':
          $ref: '#/components/responses/500'

  /mentees/{id_mentee}/feedbacks:
    get: 
      tags:
        - Mentee
      summary: ini summary
      description: ini desc
      security:
        - JWTAuth: []
      parameters:  
        - name: id_mentee
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
                  id_mentee: 1
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
                id_user: 1
                id_mentee: 1
                id_status: 1
      responses:
        '201':
          description: Register berhasil.
          content:
            application/json:
              schema:
                type: object
                example:
                  status: "success"
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

  /feedbacks/{id_feedback}:
    put:
      tags:
        - Feedback
      summary: "x"
      description: "edit data feedback"
      security:
        - JWTAuth: []
      parameters:
        - name: id_feedback
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
        - name: id_feedback
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
                  status: "success"
                  message: "success delete feedback"
        '404':
          $ref: '#/components/responses/404'
        '500':
          $ref: '#/components/responses/500'

  /teams:
    get:
      tags:
      - Team
      summary: List all teams
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
                    message: "-",
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
    
    post:
      tags:
        - Team
      summary: Create a new team
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
          description: Team Not Found
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Team'
              
              examples:
                teamData:
                  description: "Team Not Found"
                  value: {
                    status: false,
                    message: "Data not found"
                  }
    put:
      tags:
      - Team
      summary: Update a team by ID
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
          description: The updated team not found
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Team'
              
              examples:
                teamData:
                  description: "Team Not Found"
                  value: {
                    status: false,
                    message: "Data not found"
                  }
              

    delete:
      tags:
      - Team
      summary: Delete a team by ID
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

  ####################################
  #  CLASS
  ####################################

  /classes:
    get:
      tags:
        - Class
      summary: List all Class
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
    
    post:
      tags:
        - Class
      summary: Create a new class
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
        '404':
          description: Class Not Found
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Class'
              
              examples:
                teamData:
                  description: "Class Not Found"
                  value: {
                    status: false,
                    message: "Data not found"
                  }
    put:
      tags:
      - Class
      summary: Update a class by ID
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
          
        '404':
          description: The updated class not found
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Class'
              
              examples:
                teamData:
                  description: "Class Not Found"
                  value: {
                    status: false,
                    message: "Data not found"
                  }
              
    delete:
      tags:
      - Class
      summary: Delete a team by ID
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

  ####################################
  #  REQUESTED FEEDBACK
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
        phone:
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

    ResponseWithoutData:
      type: object
      properties:
        status:
          type: string
        message:
          type: string
 
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

  responses:
    '404':
      description: Not Found
      content:
        application/json:
          schema:
            type: object
            $ref: '#/components/schemas/ResponseWithoutData'
          example:
            status: "failed"
            message: "data not found"


    '500':
      description: Internal Server Error
      content:
        application/json:
          schema:
            type: object
            $ref: '#/components/schemas/ResponseWithoutData'
          example:
            status: "failed"
            message: "internal server error"


`

var SwaggerInfo = &swag.Spec{
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

func InitSwagger() {
	SwaggerInfo.Title = "API E-Absensi"
	SwaggerInfo.Description = "MIG Software Engineer"
	SwaggerInfo.Version = "1.0"
	SwaggerInfo.Host = "https://eabsensi-findryan.herokuapp.com"
	SwaggerInfo.Schemes = []string{"http", "https"}
	SwaggerInfo.BasePath = "/"
}