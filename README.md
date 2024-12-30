Get started by customizing your environment (defined in the .idx/dev.nix file) with the tools and IDE extensions you'll need for your project!

Learn more at https://developers.google.com/idx/guides/customize-idx-env

## Developing the Backend Application 

This section outlines the steps for building and running the backend application using standard Go compilation and using Docker.

### Standard Go Compilation

1. **Navigate to the backend directory:**

    ```bash
    cd backend
    ```

2. **run the compile command**

    ```bash
    $go run main.go
    ```

3. Test the app using curl command:

    ```bash
    $ curl localhost:8080
    ```

## Developing the Frontend Application

The application is build using Svelte 5 and material design.

1.  **Navigate to the frontend directory:**

    ```bash
    cd frontend
    ```

2.  **Install dependencies:**

    ```bash
    npm install
    ```
3.  **Update the CSS:**

    ```bash
    npm run prepare
    ```

4.  **Start the development server:**

    ```bash
    npm run dev
    ```

### Material design

The whole page is built on top of material design. we use those projects:

- [Material Design Icons](https://pictogrammers.com/library/mdi/)
- [Svelte Material UI](sveltematerialui.com)