import React, { useState } from 'react'
import './App.css'

interface Todo {
  id: number;
  title: string;
  done: boolean;
}

let nextId = 0;

function App() {
  const [title, setTitle] = useState<string>("");
  const [todos, setTodos] = useState<Todo[]>([]);

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const addTitle = title.trim();
    if (addTitle === '') return;
    const addTodo: Todo = { id: nextId++, title: addTitle, done: false };
    setTodos((prev => [...prev, addTodo]));
    setTitle("");
  }
  return (
    <>
      <h1>Todoリスト</h1>
      <h2>新規作成</h2>
      <form onSubmit={handleSubmit}>
        <input type="text" value={title} onChange={e => setTitle(e.target.value)} />
        <button type='submit'>追加</button>
      </form>
    </>
  )
}

export default App
