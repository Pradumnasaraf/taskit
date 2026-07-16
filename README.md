## TaskIt

It is a simple task manager that allows you to create, edit, delete and list tasks from the command line. It is built with Go and Cobra.

![TaskIt demo GIF](https://user-images.githubusercontent.com/51878265/224398202-84823ca4-ba4d-4ae4-915a-9f9f42ee50d0.gif)

### ⭐️ Features

- Adding tasks
- Deleting a task
- Updating tasks
- Completing tasks
- Deleting all tasks
- Listing all tasks

### ❗️ Prerequisites

- Have Go installed on your machine and have the basic knowledge.

### ⚙️ Installation

#### Local

At the root og the folder run the following command to download the dependencies:

```bash
go mod download
```

Then, we need to build the binary:

```bash
go build .
```

Finally, we can run the executable:

```bash
./taskit
```

### 📝 Usage

We can perform tasks like add, delete, update and list tasks. We can also delete all tasks.

**Note**: When you add task for the 1st time it will create a `tasks.json` file and all the commands will update chnages to that only. So even to you exit the terminal and come back again, the tasks will be there.

```
Usage:
  ./taskit [flags]
  ./taskit [command]

Available Commands:
  add         Add a new task
  delete      Delete a task by ID
  deleteall   Delete all tasks
  done        Mark a task as done
  help        Help about any command
  list        List all tasks
  update      Update a task by ID

Flags:
  -h, --help   help for taskit
```

### 📜 License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details

### 🛡 Security

If you discover a security vulnerability within this project, please check the [SECURITY](SECURITY.md) for more information.
