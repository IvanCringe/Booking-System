import type { Metadata } from "next";
import { AuthForm } from "@/features/auth/components/auth-form";
import { AuthShell } from "@/features/auth/components/auth-shell";

export const metadata: Metadata = {
  title: "Восстановление пароля",
};

export default function ForgotPasswordPage() {
  return (
    <AuthShell mode="forgot">
      <AuthForm mode="forgot" />
    </AuthShell>
  );
}
