# Stock Management System REST API with Clean Architecture

This project aims to develop a REST API for stock management using Golang and MongoDB, adhering to a clean architecture approach. The project is easily runnable using Docker and Docker Compose.

## Installation

1. Clone the project:

    ```bash
    git clone https://github.com/mehmet-dogru/StockManagmentSystem.git
    ```

2. Navigate to the project directory:

    ```bash
    cd StockManagmentSystem
    ```

3. Run the project using Docker Compose:

    ```bash
    docker-compose up --build
    ```

This command starts the MongoDB database and the Golang REST API server.

## Usage

You can access the API using an API client (e.g., Postman or cURL). The fundamental endpoints of the API are:

- **POST /api/v1/users/register**: Creates a new user registration.
- **POST /api/v1/users/login**: Logs in with an existing user.
- **GET /api/v1/users/account**: Provides access to the user's account information.
- **POST /api/v1/forms**: Creates a new form.
- **GET /api/v1/forms**: Retrieves all forms with pagination.
- **GET /api/v1/forms/:id**: Retrieves details of a specific form. (Replace ":id" with the form ID.)
- **PUT /api/v1/forms/:id**: Updates a specific form.
- **DELETE /api/v1/forms/:id**: Deletes a specific form.
- **POST /api/v1/forms/:id/stocks**: Adds a new stock to a form.
- **GET /api/v1/forms/:id/stocks**: Retrieves all stocks of a form.
- **GET /api/v1/forms/:id/stocks/:stock_id**: Retrieves details of a specific stock in a form. (Replace ":stock_id" with the stock ID.)
- **PUT /api/v1/forms/:id/stocks/:stock_id**: Updates a specific stock in a form.
- **DELETE /api/v1/forms/:id/stocks/:stock_id**: Deletes a specific stock from a form.

## Architecture

This project follows the principles of Clean Architecture. The core components include:

- **Domain**: The layer containing fundamental rules and data structures representing the business logic.
- **Use Cases**: The layer implementing application functionality and business workflows.
- **Interfaces**: The layer containing external components for interaction, such as HTTP Server and Database.
- **Frameworks & Drivers**: The layer directly interacting with external components and connecting lower layers.

## Contribution

If you wish to contribute to this project, please discuss your changes by opening an issue before submitting a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
