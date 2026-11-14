import Link from "next/link";

export function BrandLogo({ inverse = false }: { inverse?: boolean }) {
  return (
    <Link
      className={`brand-logo${inverse ? " brand-logo--inverse" : ""}`}
      href="/login"
      aria-label="На главную страницу Orbit"
    >
      <span className="brand-logo__mark" aria-hidden="true">
        <svg viewBox="0 0 32 32" fill="none">
          <path
            d="M16 4.5c5.6 0 10.2 4.4 10.2 9.9S21.6 24.3 16 24.3 5.8 19.9 5.8 14.4 10.4 4.5 16 4.5Z"
            stroke="currentColor"
            strokeWidth="1.8"
            opacity=".9"
          />
          <path
            d="M8.2 20.6c3.1-5.8 8.2-9.2 15.4-10.1M8.6 10.1c4.4 5.2 9.4 8 15.1 8.4"
            stroke="currentColor"
            strokeWidth="1.8"
            strokeLinecap="round"
          />
          <circle cx="23.8" cy="9" r="2.1" fill="currentColor" />
        </svg>
      </span>
      <span className="brand-logo__word">orbit<span>.</span></span>
    </Link>
  );
}
