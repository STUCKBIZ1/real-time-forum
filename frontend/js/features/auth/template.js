export function registerTemplate(){
    return `
    <div class="auth-container">
        <div class="auth-card">
            <h1 class="auth-title">
                Create Account
            </h1>
            <p class="auth-subtitle">
                Join our community today
            </p>
            <form class="auth-form" id="register-form">
                <div class="input-group">

                    <label for="firstname">
                        Firstname
                    </label>

                    <input
                    class="auth-input"
                    id="firstname"
                    type="text"
                    placeholder="John"
                    required>
                </div>
                <div class="input-group">

                    <label for="lastname">
                        Lastname
                    </label>
                    <input
                    class="auth-input"
                    id="lastname"
                    type="text"
                    placeholder="Doe"
                    required>
                </div>
                <div class="input-group">

                    <label for="nickname">
                        Nickname
                    </label>
                    <input
                    class="auth-input"
                    id="nickname"
                    type="text"
                    placeholder="john123"
                    required>
                </div>
                <div class="input-group">
                    <label for="email">
                        Email
                    </label>
                    <input
                    class="auth-input"
                    id="email"
                    type="email"
                    placeholder="example@gmail.com"
                    autocomplete="email"
                    required>
                </div>
                <div class="input-group">
                    <label for="password">
                        Password
                    </label>
                    <div class="password-wrapper">
                        <input
                        class="auth-input"
                        id="password"
                        type="password"
                        placeholder="********"
                        required>
                        <span class="password-toggle">
                            👁
                        </span>
                    </div>
                </div>
                <div class="input-group">
                    <label for="age">
                        Age
                    </label>
                    <input
                    class="auth-input"
                    id="age"
                    type="number"
                    min="13"
                    required>
                </div>
                <div class="input-group">
                    <label for="gender">
                        Gender
                    </label>
                    <select
                    class="auth-select"
                    id="gender">
                        <option value="male">
                            Male
                        </option>
                        <option value="female">
                            Female
                        </option>
                    </select>
                </div>
                <div id="register-message"></div>
                <button
                class="auth-btn"
                type="submit">
                    Register
                </button>
            </form>
            <p class="auth-footer">
                Already have an account?
                <span class="auth-link" id="login-link">
                    Login
                </span>
            </p>
        </div>
    </div>
    `;
}
export function loginTemplate(){
    return `
    <div class="auth-container">
        <div class="auth-card">
            <h1 class="auth-title">
                Welcome Back
            </h1>
            <p class="auth-subtitle">
                Login to your account
            </p>
            <form class="auth-form" id="login-form">
                <div class="input-group">
                    <label for="login-email">
                        Email
                    </label>
                    <input
                    class="auth-input"
                    id="login-email"
                    type="email"
                    placeholder="example@gmail.com"
                    autocomplete="email"
                    required>
                </div>
                <div class="input-group">
                    <label for="login-password">
                        Password
                    </label>
                    <div class="password-wrapper">
                        <input
                        class="auth-input"
                        id="login-password"
                        type="password"
                        placeholder="********"
                        required>
                        <span class="password-toggle">
                            👁
                        </span>
                    </div>
                </div>
                <div id="login-message"></div>
                <button
                class="auth-btn"
                type="submit">
                    Login
                </button>
            </form>
            <p class="auth-footer">
                Don't have an account?
                <span class="auth-link" id="register-link">
                    Register
             </span>
            </p>
        </div>
    </div>
    `;

}