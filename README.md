# Folder Tree Application

![](./src/assets/screenshot.png)

This is a simple application that displays a folder structure using a tree view with vue as frontend and golang as backend. Users can click on folders to expand them and view their contents (including files and subfolders).

## Table of Contents

- [Prerequisite](#prerequisite)
- [Installation](#installation)
- [Usage](#usage)
- [Components](#components)
  - [FolderTree](#foldertree)
  - [FolderContent](#foldercontent)
- [API Endpoints](#api-endpoints)
- [Styling](#styling)

## Prerequisite
1. Node.js installed on your system.
2. npm (Node Package Manager) installed on your system.
3. golang and postgreSQL installed on your system.
4. database migration cli `migrate` https://github.com/golang-migrate/migrate
4. A text editor or IDE of your choice.

## Installation

1. Clone this repository:
   ```bash
   git clone <repository-url>
   cd my-vue-app
   ```

2. Download dependencies for the backend:
    ```bash
    cd src/backend/folder-tree && go mod vendor
    ```

3. Migrate the database:
   ```bash
   migrate -path db/migrations -database postgres://user:password@localhost:5432/folder_tree_db migrate up
   ```
4. Run the backend server, and it will start running on 8081:
    ```bash
    go run cmd/main.go
    ```
5. Install the required front end dependencies:
   ```bash
   cd ../../.. && npm install
   ```

3. Run the development server:
   ```bash
   npm run dev
   ```

4. Open your browser and navigate to `http://localhost:5173`.

## Usage

- The left panel (`FolderTree`) displays a hierarchical structure of folders.
- Clicking on a folder will load its contents (files and subfolders) in the right panel (`FolderContent`).

## Components

### FolderTree

- **File**: `FolderTree.vue`
- **Description**: This component fetches the folder structure from an API and displays it as a tree.
- **Props**:
  - None
- **Events**:
  - `update:selected-folder`: Emits the currently selected folder.

#### Functions

- `fetchRootFolders()`: Fetches and displays the root-level folders.
- `loadChildren(parentId, parentNode)`: Loads and displays the child folders for a specific parent folder.

### FolderContent

- **File**: `FolderContent.vue`
- **Description**: Displays the contents of the currently selected folder, including files and subfolders.
- **Props**:
  - `selectedFolder`: An object representing the selected folder.
- **Events**:
  - `update:selected-folder`: Emits the selected folder when children are loaded.

#### Functions

- `clickHandler(file)`: Handles clicks on files/folders, triggering the loading of their children.

## API Endpoints

The application makes use of the following API endpoint to fetch folder data:

- **GET** `http://localhost:8081/folder/tree?id={folder_id}&depth={depth}`

  - `folder_id`: The ID of the folder to fetch.
  - `depth`: The depth of the folder tree to fetch.
