document.addEventListener("DOMContentLoaded", function () {
    // Menu Toggle
    const menuToggle = document.getElementById("menu-toggle");
    if (menuToggle) {
        menuToggle.addEventListener("click", function (e) {
            e.preventDefault();
            document.body.classList.toggle("toggled");
        });
    }

    // Dark Mode Toggle
    const darkModeToggle = document.getElementById("darkModeToggle");
    if (darkModeToggle) {
        darkModeToggle.addEventListener("click", function () {
            const htmlEl = document.documentElement;
            const icon = this.querySelector("i");
            if (htmlEl.getAttribute("data-bs-theme") === "light") {
                htmlEl.setAttribute("data-bs-theme", "dark");
                icon.classList.remove("bi-moon-stars");
                icon.classList.add("bi-sun");
            } else {
                htmlEl.setAttribute("data-bs-theme", "light");
                icon.classList.remove("bi-sun");
                icon.classList.add("bi-moon-stars");
            }
        });
    }
});
