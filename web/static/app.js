function showResult(element, message, type = "") {
    if (!element) {
        return;
    }

    element.textContent = message;

    element.classList.remove(
        "result--success",
        "result--error"
    );

    if (type === "success") {
        element.classList.add("result--success");
    }

    if (type === "error") {
        element.classList.add("result--error");
    }
}


function setButtonLoading(
    button,
    loading,
    loadingText = "Загрузка..."
) {
    if (!button) {
        return;
    }

    if (loading) {
        button.dataset.originalText = button.textContent;
        button.disabled = true;
        button.classList.add("button--loading");
        button.textContent = loadingText;

        return;
    }

    button.disabled = false;
    button.classList.remove("button--loading");

    if (button.dataset.originalText) {
        button.textContent = button.dataset.originalText;
        delete button.dataset.originalText;
    }
}


// Главная страница

const goToRegisterButton =
    document.getElementById("goToRegisterBtn");

if (goToRegisterButton) {
    goToRegisterButton.addEventListener("click", () => {
        window.location.href = "/register";
    });
}


// Регистрация

const registerForm =
    document.getElementById("registerForm");

if (registerForm) {
    registerForm.addEventListener(
        "submit",
        async (event) => {
            event.preventDefault();

            const login =
                document.getElementById("registerLogin").value;

            const password =
                document.getElementById("registerPassword").value;

            const result =
                document.getElementById("registerResult");

            const button =
                registerForm.querySelector(
                    'button[type="submit"]'
                );

            showResult(result, "");

            setButtonLoading(
                button,
                true,
                "Регистрация..."
            );

            try {
                const response = await fetch(
                    "/api/register",
                    {
                        method: "POST",

                        headers: {
                            "Content-Type": "application/json",
                        },

                        body: JSON.stringify({
                            login: login,
                            password: password,
                        }),
                    }
                );

                const data = await response.json();

                if (!response.ok) {
                    showResult(
                        result,
                        data.message,
                        "error"
                    );

                    return;
                }

                showResult(
                    result,
                    data.message,
                    "success"
                );
            } catch (error) {
                console.error(
                    "Ошибка запроса регистрации:",
                    error
                );

                showResult(
                    result,
                    "Не удалось связаться с сервером.",
                    "error"
                );
            } finally {
                setButtonLoading(button, false);
            }
        }
    );
}


// Вход

const loginForm =
    document.getElementById("loginForm");

if (loginForm) {
    loginForm.addEventListener(
        "submit",
        async (event) => {
            event.preventDefault();

            const login =
                document.getElementById("loginLogin").value;

            const password =
                document.getElementById("loginPassword").value;

            const result =
                document.getElementById("loginResult");

            const button =
                loginForm.querySelector(
                    'button[type="submit"]'
                );

            showResult(result, "");

            setButtonLoading(
                button,
                true,
                "Входим..."
            );

            try {
                const response = await fetch(
                    "/api/login",
                    {
                        method: "POST",

                        headers: {
                            "Content-Type": "application/json",
                        },

                        body: JSON.stringify({
                            login: login,
                            password: password,
                        }),
                    }
                );

                const data = await response.json();

                if (!response.ok) {
                    showResult(
                        result,
                        data.message,
                        "error"
                    );

                    return;
                }

                showResult(
                    result,
                    data.message,
                    "success"
                );

                window.location.href = "/cabinet";
            } catch (error) {
                console.error(
                    "Ошибка запроса входа:",
                    error
                );

                showResult(
                    result,
                    "Не удалось связаться с сервером.",
                    "error"
                );
            } finally {
                setButtonLoading(button, false);
            }
        }
    );
}


// Сообщение

const messageForm =
    document.getElementById("messageForm");

if (messageForm) {
    const messageText =
        document.getElementById("messageText");

    const messageResult =
        document.getElementById("messageResult");

    const saveButton =
        messageForm.querySelector(
            'button[type="submit"]'
        );

    const loadMessage = async () => {
        showResult(
            messageResult,
            "Загружаем сообщение..."
        );

        try {
            const response =
                await fetch("/api/message");

            if (response.status === 401) {
                window.location.href = "/login";
                return;
            }

            if (!response.ok) {
                showResult(
                    messageResult,
                    "Не удалось загрузить сообщение.",
                    "error"
                );

                return;
            }

            const data = await response.json();

            messageText.value = data.text;

            if (data.text) {
                showResult(
                    messageResult,
                    "Сообщение загружено.",
                    "success"
                );
            } else {
                showResult(
                    messageResult,
                    "Пока здесь ничего не сохранено."
                );
            }
        } catch (error) {
            console.error(
                "Ошибка загрузки сообщения:",
                error
            );

            showResult(
                messageResult,
                "Не удалось связаться с сервером.",
                "error"
            );
        }
    };


    messageForm.addEventListener(
        "submit",
        async (event) => {
            event.preventDefault();

            showResult(messageResult, "");

            setButtonLoading(
                saveButton,
                true,
                "Сохраняем..."
            );

            try {
                const response = await fetch(
                    "/api/message",
                    {
                        method: "POST",

                        headers: {
                            "Content-Type":
                                "application/json",
                        },

                        body: JSON.stringify({
                            text: messageText.value,
                        }),
                    }
                );

                const data = await response.json();

                if (response.status === 401) {
                    window.location.href = "/login";
                    return;
                }

                if (!response.ok) {
                    showResult(
                        messageResult,
                        data.message,
                        "error"
                    );

                    return;
                }

                showResult(
                    messageResult,
                    data.message,
                    "success"
                );
            } catch (error) {
                console.error(
                    "Ошибка сохранения сообщения:",
                    error
                );

                showResult(
                    messageResult,
                    "Не удалось связаться с сервером.",
                    "error"
                );
            } finally {
                setButtonLoading(
                    saveButton,
                    false
                );
            }
        }
    );

    loadMessage();
}


// Выход

const logoutButton =
    document.getElementById("logoutBtn");

if (logoutButton) {
    logoutButton.addEventListener(
        "click",
        async () => {
            setButtonLoading(
                logoutButton,
                true,
                "Выходим..."
            );

            try {
                const response = await fetch(
                    "/api/logout",
                    {
                        method: "POST",
                    }
                );

                if (!response.ok) {
                    console.error(
                        "Не удалось выполнить выход"
                    );

                    setButtonLoading(
                        logoutButton,
                        false
                    );

                    return;
                }

                window.location.href = "/";
            } catch (error) {
                console.error(
                    "Ошибка запроса выхода:",
                    error
                );

                setButtonLoading(
                    logoutButton,
                    false
                );
            }
        }
    );
}