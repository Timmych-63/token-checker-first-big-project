const goToRegisterButton = document.getElementById("goToRegisterBtn");

if (goToRegisterButton) {
    goToRegisterButton.addEventListener("click", () =>{
        window.location.href = "/register";
    });
}

const registerForm = document.getElementById("registerForm");

if (registerForm) {
    registerForm.addEventListener("submit", async (event) => {
        event.preventDefault();

        const login = document.getElementById("registerLogin").value;
        const password = document.getElementById("registerPassword").value;
        const result = document.getElementById("registerResult");

        const response = await fetch("/api/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body:  JSON.stringify({
                login: login,
                password: password,
            }),
        });

        const data = await response.json();

        if(!response.ok) {
            result.textContent = data.message;
            return;
        }        

        result.textContent = data.message;
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