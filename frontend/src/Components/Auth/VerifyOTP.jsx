import { useState } from "react";
import { useForm } from "react-hook-form";
import { Link, useNavigate, useSearchParams } from "react-router-dom";

import { verifyOTP } from "../../api/auth";

export default function VerifyOTP() {
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const userId = searchParams.get("id") || "";

  const [serverError, setServerError] = useState("");

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm({ mode: "onBlur" });

  const onSubmit = async (data) => {
    try {
      setServerError("");
      await verifyOTP(userId, data.otp);
      navigate("/auth/login", { replace: true });
    } catch (err) {
      setServerError(err?.message || "Verification failed");
    }
  };

  return (
    <div className="page">
      <div className="container">
        <form onSubmit={handleSubmit(onSubmit)} className="ui-panel pad-6 auth-card">
          <div className="row-between">
            <div>
              <div className="auth-title">Verify OTP</div>
              <div className="auth-sub">Check your email, then enter the 6 digits</div>
            </div>
            <div className="ui-tag">auth</div>
          </div>

          {serverError ? (
            <div className="alert alert-danger">
              {serverError}
            </div>
          ) : null}

          <div className="form">
            <input
              type="text"
              placeholder="000000"
              inputMode="numeric"
              className="ui-input otp-input"
              {...register("otp", {
                required: "OTP is required",
                pattern: { value: /^\d{6}$/, message: "OTP must be 6 digits" },
              })}
            />
            {errors.otp ? <div className="field-error">{errors.otp.message}</div> : null}

            {!userId ? <div className="small muted2">Missing user id in URL</div> : null}
          </div>

          <button className="ui-btn-primary full" type="submit" disabled={isSubmitting || !userId}>
            {isSubmitting ? "Verifying..." : "verify"}
          </button>

          <div className="help-row small muted">
            <div className="muted">Done?</div>
            <Link className="ui-btn-secondary" to="/auth/login">
              back to login
            </Link>
          </div>
        </form>
      </div>
    </div>
  );
}
