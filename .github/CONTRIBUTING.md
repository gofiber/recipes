# Contributing Guidelines

Thank you for considering contributing to this project! To ensure a smooth and efficient process, please follow these guidelines.

## Adding a New Example

1. **Create a Directory**: Create a new directory for your example in the root of the repository. Please do not use a "fiber" prefix in the directory name.

2. **Add a `README.md`**: Each example must include a `README.md` file in its directory. This file should contain the following:

  - **Docusaurus Metadata**: Add the following metadata at the top of the `README.md` file:
    ```markdown
    ---
    title: Your Example Title
    keywords: [keyword1, keyword2, keyword3]
    ---
    ```

    - `title`: A short and descriptive title for your example.
    - `keywords`: A list of relevant keywords (excluding "fiber").

  - **Content**: The `README.md` should provide a detailed explanation of the example, including:
    - The idea behind the example.
    - The components used in the example.
    - Instructions on how to run the example.
    - Any other relevant information.

3. **Update the Overview**: After adding your example, run the following command in the root directory to update the overview table of contents:
    ```bash
    make generate
    ```

By following these guidelines, you help maintain the quality and consistency of the project. Thank you for your contributions!
