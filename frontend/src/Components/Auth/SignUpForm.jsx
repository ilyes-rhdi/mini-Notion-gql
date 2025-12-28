import { useState } from "react";
import { useForm } from "react-hook-form";
import { Link, useNavigate } from "react-router-dom";

import { register as registerUser } from "../../api/auth";

function toBoolGender(value) {
  if (value === "true") return true;
  if (value === "false") return false;
  return true;
}

export default function SignUpForm() {
  const navigate = useNavigate();
  const [serverError, setServerError] = useState("");

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm({ mode: "onBlur" });

  const onSubmit = async (data) => {
    try {
      setServerError("");
      const payload = { ...data, gender: toBoolGender(data.gender) };
      const res = await registerUser(payload);
      const id = res?.userId || res?.id || res?.user?.id || res?.userID;

      if (id) {
        navigate(`/auth/verify?id=${encodeURIComponent(id)}`, { replace: true });
        return;
      }

      navigate("/auth/login", { replace: true });
    } catch (err) {
      setServerError(err?.message || "Signup failed");
    }
  };

  return (
    <div className="page">
      <div className="container">
        <form onSubmit={handleSubmit(onSubmit)} className="ui-panel pad-6 auth-card">
          <div className="row-between">
            <div>
              <div className="auth-title">Create account</div>
              <div className="auth-sub">Then you can create workspaces + pages</div>
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
              <input className="ui-input" placeholder="Name" {...register("name", { required: "Name is required" })} />
              {errors.name ? <div className="field-error">{errors.name.message}</div> : null}
            </div>

            <div>
              <input
                type="email"
                className="ui-input"
                placeholder="Email"
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
                className="ui-input"
                placeholder="Password"
                {...register("password", {
                  required: "Password is required",
                  minLength: { value: 6, message: "Min 6 characters" },
                })}
              />
              {errors.password ? <div className="field-error">{errors.password.message}</div> : null}
            </div>

            <div>
              <select className="ui-select" {...register("gender", { required: "Gender is required" })} defaultValue="">
                <option value="" disabled>
                  Gender
                </option>
                <option value="true">Male</option>
                <option value="false">Female</option>
              </select>
              {errors.gender ? <div className="field-error">{errors.gender.message}</div> : null}
            </div>
          </div>

          <button className="ui-btn-primary full" type="submit" disabled={isSubmitting}>
            {isSubmitting ? "Creating..." : "create"}
          </button>

          <div className="help-row small muted">
            <div className="muted">Already have an account?</div>
            <Link className="ui-btn-secondary" to="/auth/login">
              login
            </Link>
          </div>
        </form>
      </div>
    </div>
  );
}
