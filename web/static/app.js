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
            console.error(
                "Ошибка запроса регистрации:",
                error,
            );

            result.textContent =
                "Не удалось связаться с сервером.";
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

            if (!response.ok) {
                return;
            }

            window.location.href = "/cabinet";
        } catch (error) {
            console.error(
                "Ошибка запроса входа:",
                error,
            );

            result.textContent =
                "Не удалось связаться с сервером.";
        }
    });
}


// Сообщение

const messageForm = document.getElementById("messageForm");

if (messageForm) {
    const messageText = document.getElementById("messageText");
    const messageResult = document.getElementById("messageResult");

    const loadMessage = async () => {
        try {
            const response = await fetch("/api/message");

            if (response.status === 401) {
                window.location.href = "/login";
                return;
            }

            if (!response.ok) {
                messageResult.textContent =
                    "Не удалось загрузить сообщение.";
                return;
            }

            const data = await response.json();

            messageText.value = data.text;
        } catch (error) {
            console.error(
                "Ошибка загрузки сообщения:",
                error,
            );

            messageResult.textContent =
                "Не удалось связаться с сервером.";
        }
    };

    messageForm.addEventListener("submit", async (event) => {
        event.preventDefault();

        try {
            const response = await fetch("/api/message", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    text: messageText.value,
                }),
            });

            const data = await response.json();

            if (response.status === 401) {
                window.location.href = "/login";
                return;
            }

            messageResult.textContent = data.message;
        } catch (error) {
            console.error(
                "Ошибка сохранения сообщения:",
                error,
            );

            messageResult.textContent =
                "Не удалось связаться с сервером.";
        }
    });

    loadMessage();
}


// Выход

const logoutButton = document.getElementById("logoutBtn");

if (logoutButton) {
    logoutButton.addEventListener("click", () => {
        window.location.href = "/";
    });
}