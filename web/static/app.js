const goToRegisterButton = document.getElementById("goToRegisterBtn");

if (goToRegisterButton) {
    goToRegisterButton.addEventListener("click", () => {
        window.location.href = "/register";
    });
}


// Регистрация

const registerForm = document.getElementById("registerForm");

if (registerForm) {
    registerForm.addEventListener("submit", async (event) => {
        event.preventDefault();

        const login = document.getElementById("registerLogin").value;
        const password = document.getElementById("registerPassword").value;
        const result = document.getElementById("registerResult");

        try {
            const response = await fetch("/api/register", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    login: login,
                    password: password,
                }),
            });

            const data = await response.json();

            result.textContent = data.message;
        } catch (error) {
            console.error("Ошибка запроса регистрации:", error);

            result.textContent = "Не удалось связаться с сервером.";
        }
    });
}


// Вход

const loginForm = document.getElementById("loginForm");

if (loginForm) {
    loginForm.addEventListener("submit", async (event) => {
        event.preventDefault();

        const login = document.getElementById("loginLogin").value;
        const password = document.getElementById("loginPassword").value;
        const result = document.getElementById("loginResult");

        try {
            const response = await fetch("/api/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    login: login,
                    password: password,
                }),
            });

            const data = await response.json();

            result.textContent = data.message;
        } catch (error) {
            console.error("Ошибка запроса входа:", error);

            result.textContent = "Не удалось связаться с сервером.";
        }
    });
}


// Сообщение

const messageForm = document.getElementById("messageForm");

if (messageForm) {
    messageForm.addEventListener("submit", (event) => {
        event.preventDefault();

        const message = document.getElementById("messageText").value;
        const result = document.getElementById("messageResult");

        console.log("Сообщение:", message);

        result.textContent =
            "Пока сообщение просто выводится в консоль.";
    });
}


// Выход

const logoutButton = document.getElementById("logoutBtn");

if (logoutButton) {
    logoutButton.addEventListener("click", () => {
        console.log("Выход из кабинета");

        window.location.href = "/";
    });
}