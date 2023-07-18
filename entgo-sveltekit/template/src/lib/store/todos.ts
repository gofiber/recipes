import { writable } from 'svelte/store';
//import { PUBLIC_API_URL as uri } from '$env/static/public'

/**
 * Represents the type of a todo item.
 */
export type TodoType = {
    id?: string;
    content: string;
    completed?: boolean;
}

export const loading = writable(false)

export const error = writable(false)

/**
 * Create a writable store with an empty array as the initial value
 */
export const todos = writable<TodoType[]>([]);

/**
 * Adds a new todo item to the store.
 */
export const add = async (todo: TodoType) => {
    try {
        const response = await fetch("/api/v1/todo/create", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(todo)
        });
        const result = await response.json() as TodoType;
        todos.update(existingTodos => [...existingTodos, result]);
    } catch (e) {
        error.set(true);
    }
};

/**
 * Removes a todo item from the store based on its ID.
 */
export const remove = async (id: string | undefined) => {
    if (!id) {
        return;
    }
    try {
        await fetch(`/api/v1/todo/delete/${id}`, {
            method: 'DELETE'
        });
        todos.update(existingTodos => existingTodos.filter(todo => todo.id !== id));
    } catch (e) {
        error.set(true);
    }
};

/*
* Updates a todo item in the store based on its ID.
*/
export const update = async (id: string | undefined, updatedTodo: TodoType) => {
    if (!id) {
        return;
    }
    try {
        const response = await fetch(`/api/v1/todo/update/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(updatedTodo)
        });
        const result = await response.json() as TodoType;
        todos.update(existingTodos => {
            return existingTodos.map(todo => {
                if (todo.id === id) {
                    return result;
                }
                return todo;
            });
        });
    } catch (error) {
        console.error(error); // Log the error to the console
    }
};