import type { ReactNode, SVGProps } from "react";

type IconName =
  | "arrow-left"
  | "arrow-right"
  | "check"
  | "check-circle"
  | "eye"
  | "eye-off"
  | "lock"
  | "mail"
  | "sparkles"
  | "user";

type IconProps = SVGProps<SVGSVGElement> & {
  name: IconName;
  size?: number;
};

const iconPaths: Record<IconName, ReactNode> = {
  "arrow-left": <path d="m14.5 18-6-6 6-6M8.5 12h11" />,
  "arrow-right": <path d="M4.5 12h15m-6-6 6 6-6 6" />,
  check: <path d="m5 12 4.2 4.2L19 6.5" />,
  "check-circle": (
    <>
      <circle cx="12" cy="12" r="9" />
      <path d="m8 12.2 2.5 2.5L16.5 9" />
    </>
  ),
  eye: (
    <>
      <path d="M2.5 12s3.4-6 9.5-6 9.5 6 9.5 6-3.4 6-9.5 6-9.5-6-9.5-6Z" />
      <circle cx="12" cy="12" r="2.6" />
    </>
  ),
  "eye-off": (
    <>
      <path d="m3 3 18 18M10.6 6.2A10.7 10.7 0 0 1 12 6c6.1 0 9.5 6 9.5 6a16 16 0 0 1-3.1 3.7M6.2 6.8C3.8 8.4 2.5 12 2.5 12s3.4 6 9.5 6c1 0 1.9-.2 2.7-.5" />
      <path d="M9.9 9.9a3 3 0 0 0 4.2 4.2" />
    </>
  ),
  lock: (
    <>
      <rect x="4.5" y="10" width="15" height="11" rx="2.2" />
      <path d="M8 10V7.5a4 4 0 0 1 8 0V10m-4 4v3" />
    </>
  ),
  mail: (
    <>
      <rect x="3" y="5" width="18" height="14" rx="2.2" />
      <path d="m4 7 8 6 8-6" />
    </>
  ),
  sparkles: (
    <>
      <path d="m12 3 1.8 5.2L19 10l-5.2 1.8L12 17l-1.8-5.2L5 10l5.2-1.8L12 3Z" />
      <path d="m19 15 .9 2.1L22 18l-2.1.9L19 21l-.9-2.1L16 18l2.1-.9L19 15Z" />
    </>
  ),
  user: (
    <>
      <circle cx="12" cy="8" r="3.5" />
      <path d="M5 20a7 7 0 0 1 14 0" />
    </>
  ),
};

export function Icon({ name, size = 18, ...props }: IconProps) {
  return (
    <svg
      aria-hidden="true"
      width={size}
      height={size}
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="1.7"
      strokeLinecap="round"
      strokeLinejoin="round"
      focusable="false"
      {...props}
    >
      {iconPaths[name]}
    </svg>
  );
}
