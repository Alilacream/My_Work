async function loadTodos() {
    const res = await fetch('/todo');
    const todos = await res.json();
    const tbody = document.querySelector('#todo-table tbody');
    tbody.innerHTML = '';
    todos.forEach(todo => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>${todo.id}</td>
            <td class="${todo.done ? 'done' : ''}">${todo.purpose}</td>
            <td><input type="checkbox" ${todo.done ? 'checked' : ''} onchange="updateDone(${todo.id}, this.checked)"></td>
            <td><button onclick="deleteTodo(${todo.id})">Delete</button></td>
        `;
        tbody.appendChild(tr);
    });
}
async function updateDone(id, done) {
    await fetch(`/todo/${id}`, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ done })
    });
    loadTodos();
}
async function deleteTodo(id) {
    await fetch(`/todo/${id}`, { method: 'DELETE' });
    loadTodos();
}
document.getElementById('add-form').onsubmit = async e => {
    e.preventDefault();
    const purpose = document.getElementById('purpose').value;
    await fetch('/todo', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ purpose })
    });
    document.getElementById('purpose').value = '';
    loadTodos();
};
loadTodos();
