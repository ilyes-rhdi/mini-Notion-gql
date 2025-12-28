import { useState } from "react";
import { useForm } from "react-hook-form";
import { Link, useNavigate } from "react-router-dom";

import { login } from "../../api/auth";
import useAuth from "../../hooks/useAuth";

export default function LoginForm() {
  const navigate = useNavigate();
  const { setToken } = useAuth();
  const [serverError, setServerError] = useState("");

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm({ mode: "onBlur" });

  const onSubmit = async (data) => {
    try {
      setServerError("");
      const res = await login(data);
      const token = res?.token || res?.accessToken || res?.jwt || "";
      if (!token) {
        setServerError("No token received from server");
        return;
      }
      setToken(token);
      navigate("/app", { replace: true });
    } catch (err) {
      setServerError(err?.message || "Login failed");
    }
  };

  return (
    <div className="page">
      <div className="container">
        <form onSubmit={handleSubmit(onSubmit)} className="ui-panel pad-6 auth-card">
          <div className="row-between">
            <div>
              <div className="auth-title">Login</div>
              <div className="auth-sub">Use your email + password</div>
            </div>
            <div className="ui-tag">auth</div>
          </div>

          {serverError ? (
            <div className="alert alert-danger">
              {serverError}
            </div>
          ) : null}

          <div className="form">
            <div>
              <input
                type="email"
                placeholder="Email"
                className="ui-input"
                {...register("email", {
                  required: "Email is required",
                  pattern: {
                    value: /^[^\s@]+@[^\s@]+\.[^\s@]+$/,
                    message: "Invalid email",
                  },
                })}
              />
              {errors.email ? <div className="field-error">{errors.email.message}</div> : null}
            </div>

            <div>
              <input
                type="password"
                placeholder="Password"
                className="ui-input"
                {...register("password", {
                  required: "Password is required",
                  minLength: { value: 6, message: "Min 6 characters" },
                })}
              />
              {errors.password ? (
                <div className="field-error">{errors.password.message}</div>
              ) : null}
            </div>
          </div>

          <button className="ui-btn-primary full" type="submit" disabled={isSubmitting}>
            {isSubmitting ? "Logging in..." : "login"}
          </button>

          <div className="help-row small muted">
            <div className="muted">No account?</div>
            <Link className="ui-btn-secondary" to="/auth/signup">
              create account
            </Link>
          </div>
        </form>
      </div>
    </div>
  );
}
