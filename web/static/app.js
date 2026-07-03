const goToRegisterButton = document.getElementById("goToRegisterBtn");

if (goToRegisterButton) {
    goToRegisterButton.addEventListener("click", () =>{
        window.location.href = "/register";
    });
}

const registerForm = document.getElementById("registerForm");

if (registerForm) {
    registerForm.addEventListener("submit", (event) => {
        event.preventDefault();

        const login = document.getElementById("registerLogin").value;
        const password = document.getElementById("registerPassword").value;

        console.log("Регистрация:");
        console.log("Логин:", login);
        console.log("Пароль:", password);

        const result = document.getElementById("registerResult");
        result.textContent = "Пока данные просто выводятся в консоль.";
    });
}

const messageForm = document.getElementById("messageForm");

if(messageForm) {
    messageForm.addEventListener("submit", (event) => {
        event.preventDefault();

        const message = document.getElementById("messageText").value;

        console.log("Сообщение:");
        console.log(message);

        const result = document.getElementById("messageResult");
        result.textContent = "Пока сообщение просто выводится в консоль.";
    });
}

const logoutButton = document.getElementById("logoutBtn");

    if (logoutButton) {
        logoutButton.addEventListener("click", () => {
            console.log("Выход из кабинета");

            window.location.href = "/";
        });
    }