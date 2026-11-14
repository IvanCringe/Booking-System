"use client";

import Link from "next/link";
import { useState, type FormEvent } from "react";
import { Icon } from "./icon";

type AuthMode = "login" | "signup" | "forgot";
type FieldName = "name" | "email" | "password";
type FormValues = Record<FieldName, string>;
type FormErrors = Partial<Record<FieldName | "terms", string>>;
type Feedback = {
  tone: "info" | "success";
  message: string;
} | null;

const initialValues: FormValues = {
  name: "",
  email: "",
  password: "",
};

const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

function validate(
  mode: AuthMode,
  values: FormValues,
  acceptedTerms: boolean,
): FormErrors {
  const errors: FormErrors = {};

  if (mode === "signup" && values.name.trim().length < 2) {
    errors.name = "Введите имя — не менее двух символов.";
  }

  if (!emailPattern.test(values.email.trim())) {
    errors.email = "Введите корректный адрес электронной почты.";
  }

  if (mode !== "forgot" && values.password.length < 8) {
    errors.password = "Пароль должен содержать не менее 8 символов.";
  }

  if (mode === "signup" && !acceptedTerms) {
    errors.terms = "Чтобы создать аккаунт, подтвердите согласие с условиями.";
  }

  return errors;
}

type TextFieldProps = {
  id: string;
  label: string;
  value: string;
  placeholder: string;
  type?: "text" | "email";
  autoComplete: string;
  error?: string;
  icon: "user" | "mail";
  onChange: (value: string) => void;
};

function TextField({
  id,
  label,
  value,
  placeholder,
  type = "text",
  autoComplete,
  error,
  icon,
  onChange,
}: TextFieldProps) {
  return (
    <div className={`form-field${error ? " form-field--error" : ""}`}>
      <label className="form-field__label" htmlFor={id}>{label}</label>
      <div className="form-field__control">
        <Icon name={icon} size={18} className="form-field__icon" />
        <input
          id={id}
          name={id}
          type={type}
          value={value}
          placeholder={placeholder}
          autoComplete={autoComplete}
          onChange={(event) => onChange(event.target.value)}
          aria-invalid={Boolean(error)}
          aria-describedby={error ? `${id}-error` : undefined}
          required
        />
      </div>
      {error ? <p className="form-field__error" id={`${id}-error`}>{error}</p> : null}
    </div>
  );
}

type PasswordFieldProps = {
  value: string;
  error?: string;
  autoComplete: string;
  onChange: (value: string) => void;
};

function PasswordField({ value, error, autoComplete, onChange }: PasswordFieldProps) {
  const [visible, setVisible] = useState(false);

  return (
    <div className={`form-field${error ? " form-field--error" : ""}`}>
      <label className="form-field__label" htmlFor="password">Пароль</label>
      <div className="form-field__control">
        <Icon name="lock" size={18} className="form-field__icon" />
        <input
          id="password"
          name="password"
          type={visible ? "text" : "password"}
          value={value}
          placeholder="Введите пароль"
          autoComplete={autoComplete}
          onChange={(event) => onChange(event.target.value)}
          aria-invalid={Boolean(error)}
          aria-describedby={error ? "password-error" : undefined}
          required
        />
        <button
          className="form-field__trailing"
          type="button"
          onClick={() => setVisible((current) => !current)}
          aria-label={visible ? "Скрыть пароль" : "Показать пароль"}
          aria-pressed={visible}
        >
          <Icon name={visible ? "eye-off" : "eye"} size={18} />
        </button>
      </div>
      {error ? <p className="form-field__error" id="password-error">{error}</p> : null}
    </div>
  );
}

function GoogleMark() {
  return <span className="google-mark" aria-hidden="true">G</span>;
}

function PasswordStrength({ password }: { password: string }) {
  const checks = [
    password.length >= 8,
    /[a-z]/.test(password) && /[A-Z]/.test(password),
    /\d/.test(password) || /[^a-zA-Z0-9]/.test(password),
  ];
  const score = checks.filter(Boolean).length;
  const labels = ["Введите пароль", "Можно надёжнее", "Уже лучше", "Надёжный пароль"];

  if (!password) return null;

  return (
    <div className="password-strength" aria-live="polite">
      <div className="password-strength__bars" aria-hidden="true">
        {checks.map((passed, index) => (
          <span className={passed ? "is-filled" : ""} key={index} />
        ))}
      </div>
      <span>{labels[score]}</span>
    </div>
  );
}

export function AuthForm({ mode }: { mode: AuthMode }) {
  const [values, setValues] = useState<FormValues>(initialValues);
  const [errors, setErrors] = useState<FormErrors>({});
  const [rememberMe, setRememberMe] = useState(true);
  const [acceptedTerms, setAcceptedTerms] = useState(false);
  const [feedback, setFeedback] = useState<Feedback>(null);

  const copy = {
    login: {
      eyebrow: "СНОВА РАДЫ ВАС ВИДЕТЬ",
      title: "С возвращением",
      description: "Войдите, чтобы продолжить работу с того места, на котором остановились.",
      button: "Войти",
      autoComplete: "current-password",
    },
    signup: {
      eyebrow: "НАЧНИТЕ БЕСПЛАТНО",
      title: "Создайте аккаунт",
      description: "Соберите задачи и проекты команды в одном понятном пространстве.",
      button: "Создать аккаунт",
      autoComplete: "new-password",
    },
    forgot: {
      eyebrow: "ВОССТАНОВЛЕНИЕ ДОСТУПА",
      title: "Восстановление пароля",
      description: "Укажите почту, к которой привязан аккаунт. Мы подскажем, что делать дальше.",
      button: "Отправить инструкции",
      autoComplete: "off",
    },
  }[mode];

  function updateValue(field: FieldName, value: string) {
    setValues((current) => ({ ...current, [field]: value }));
    setErrors((current) => ({ ...current, [field]: undefined }));
    setFeedback(null);
  }

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    const nextErrors = validate(mode, values, acceptedTerms);
    setErrors(nextErrors);

    if (Object.keys(nextErrors).length > 0) {
      setFeedback({
        tone: "info",
        message: "Проверьте поля формы — некоторые данные нужно исправить.",
      });
      return;
    }

    const messages = {
      login: "Проверка формы пройдена. Настоящий вход станет доступен после подключения сервера.",
      signup: "Проверка пройдена. Создание аккаунта станет доступно после подключения сервера.",
      forgot: "Адрес почты указан верно. Отправка писем для сброса пароля пока не подключена.",
    };

    setFeedback({ tone: "success", message: messages[mode] });
  }

  function handleGoogleClick() {
    setFeedback({
      tone: "info",
      message: "Вход через Google пока не подключён. Эта кнопка станет рабочей после настройки авторизации через Google.",
    });
  }

  return (
    <div className="auth-form-wrap">
      <div className="auth-heading">
        <div className="auth-heading__eyebrow">
          <span className="auth-heading__spark"><Icon name="sparkles" size={14} /></span>
          {copy.eyebrow}
        </div>
        <h1>{copy.title}<span>.</span></h1>
        <p>{copy.description}</p>
      </div>

      {mode !== "forgot" ? (
        <>
          <button className="google-button" type="button" onClick={handleGoogleClick}>
            <GoogleMark />
            Продолжить с Google
          </button>
          <div className="form-divider"><span>или войдите по электронной почте</span></div>
        </>
      ) : (
        <div className="forgot-note">
          <span className="forgot-note__icon"><Icon name="mail" size={18} /></span>
          <p>Когда мы подключим наш любимый бэкендик, здесь появится подтверждение отправки письма для восстановления пароля.</p>
        </div>
      )}

      <form className="auth-form" onSubmit={handleSubmit} noValidate>
        {mode === "signup" ? (
          <TextField
            id="name"
            label="Ваше имя"
            value={values.name}
            placeholder="Ванек Пес ХУлиганский"
            autoComplete="name"
            icon="user"
            error={errors.name}
            onChange={(value) => updateValue("name", value)}
          />
        ) : null}

        <TextField
          id="email"
          label="Электронная почта"
          value={values.email}
          placeholder="ivan@company.ru"
          type="email"
          autoComplete="email"
          icon="mail"
          error={errors.email}
          onChange={(value) => updateValue("email", value)}
        />

        {mode !== "forgot" ? (
          <div className="auth-form__password-block">
            <div className="auth-form__password-label-row">
              {mode === "login" ? (
                <Link className="text-link" href="/forgot-password">Забыли пароль?</Link>
              ) : null}
            </div>
            <PasswordField
              value={values.password}
              autoComplete={copy.autoComplete}
              error={errors.password}
              onChange={(value) => updateValue("password", value)}
            />
            {mode === "signup" ? <PasswordStrength password={values.password} /> : null}
          </div>
        ) : null}

        {mode === "login" ? (
          <label className="check-row">
            <input
              type="checkbox"
              checked={rememberMe}
              onChange={(event) => setRememberMe(event.target.checked)}
            />
            <span className="check-row__box"><Icon name="check" size={12} /></span>
            <span>Запомнить меня</span>
          </label>
        ) : null}

        {mode === "signup" ? (
          <div className="terms-wrap">
            <label className="check-row check-row--terms">
              <input
                type="checkbox"
                checked={acceptedTerms}
                onChange={(event) => {
                  setAcceptedTerms(event.target.checked);
                  setErrors((current) => ({ ...current, terms: undefined }));
                }}
              />
              <span className="check-row__box"><Icon name="check" size={12} /></span>
              <span>Я принимаю Пользовательское соглашение и Политику конфиденциальности.</span>
            </label>
            {errors.terms ? <p className="form-field__error terms-error">{errors.terms}</p> : null}
          </div>
        ) : null}

        <button className="submit-button" type="submit">
          <span>{copy.button}</span>
          <span className="submit-button__icon"><Icon name="arrow-right" size={17} /></span>
        </button>

        {feedback ? (
          <div className={`form-feedback form-feedback--${feedback.tone}`} role="status" aria-live="polite">
            <Icon name={feedback.tone === "success" ? "check-circle" : "sparkles"} size={17} />
            <span>{feedback.message}</span>
          </div>
        ) : null}
      </form>

      <div className="auth-bottom-note">
        <Icon name="lock" size={14} />
        <span>.</span>
      </div>
    </div>
  );
}
