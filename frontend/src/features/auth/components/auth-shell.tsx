import Link from "next/link";
import type { ReactNode } from "react";
import { BrandLogo } from "./brand-logo";
import { Icon } from "./icon";

type AuthMode = "login" | "signup" | "forgot";

type AuthShellProps = {
  mode: AuthMode;
  children: ReactNode;
};

const topLinks: Record<AuthMode, { label: string; href: string; action: string }> = {
  login: {
    label: "Впервые в Orbit?",
    href: "/signup",
    action: "Создать аккаунт",
  },
  signup: {
    label: "Уже есть аккаунт?",
    href: "/login",
    action: "Войти",
  },
  forgot: {
    label: "Вспомнили пароль?",
    href: "/login",
    action: "Вернуться ко входу",
  },
};

export function AuthShell({ mode, children }: AuthShellProps) {
  const topLink = topLinks[mode];

  return (
    <main className="auth-stage">
      <div className="auth-stage__glow auth-stage__glow--one" aria-hidden="true" />
      <div className="auth-stage__glow auth-stage__glow--two" aria-hidden="true" />

      <div className="auth-frame">
        <section className="auth-main" aria-label="Вход и регистрация">
          <header className="auth-header">
            <BrandLogo />
            <div className="auth-header__switch">
              <span>{topLink.label}</span>
              <Link href={topLink.href}>{topLink.action}</Link>
            </div>
          </header>

          <div className="auth-main__content">{children}</div>

          <footer className="auth-footer">
            <span>ООО "ARTEM CELULIT"</span>
            <span className="auth-footer__secure">
              <Icon name="lock" size={13} />
              Предварительная версия рабочего пространства
            </span>
          </footer>
        </section>

        <aside className="story-panel" aria-label="О сервисе Orbit">
          <div className="story-panel__texture" aria-hidden="true" />
          <div className="story-panel__orb story-panel__orb--one" aria-hidden="true" />
          <div className="story-panel__orb story-panel__orb--two" aria-hidden="true" />

          <div className="story-panel__inner">
            <div className="story-panel__top">
              <BrandLogo inverse />
              <div className="story-panel__availability">
                <span /> Больше спокойствия в работе
              </div>
            </div>

            <div className="story-copy">
              <div className="story-copy__eyebrow">
                <Icon name="sparkles" size={14} />
                КОМАНДА НА ОДНОЙ ВОЛНЕ
              </div>
              <h2>
                Освободите время для
                <br />
                <span>действительно важного.</span>
              </h2>
              <p>
                Объедините проекты, команду и повседневные задачи в одном понятном
                пространстве. Меньше суеты — больше результата.
              </p>
            </div>

            <div className="workspace-preview" aria-label="Пример рабочего пространства Orbit">
              <div className="workspace-preview__topbar">
                <div className="workspace-preview__dots" aria-hidden="true">
                  <i />
                  <i />
                  <i />
                </div>
                <span>orbit / пространство</span>
                <div className="workspace-preview__avatar">А</div>
              </div>

              <div className="workspace-preview__body">
                <div className="workspace-preview__heading">
                  <div>
                    <span className="workspace-preview__overline">ОБЗОР ЗА НЕДЕЛЮ</span>
                    <strong>Доброе утро, команда</strong>
                  </div>
                  <span className="workspace-preview__period">Эта неделя⌄</span>
                </div>

                <div className="workspace-preview__metrics">
                  <div className="workspace-metric">
                    <span>Выполнено задач</span>
                    <strong>24 <small>/ 32</small></strong>
                    <div className="workspace-metric__bar"><i /></div>
                  </div>
                  <div className="workspace-metric workspace-metric--violet">
                    <span>Фокус команды</span>
                    <strong>86<small>%</small></strong>
                    <div className="workspace-metric__bar"><i /></div>
                  </div>
                </div>

                <div className="workspace-preview__list-head">
                  <strong>Главные задачи на сегодня</strong>
                  <span>Все задачи</span>
                </div>
                <div className="workspace-task">
                  <span className="workspace-task__check"><Icon name="check" size={11} /></span>
                  <span>Согласовать направление продукта</span>
                  <small>Готово</small>
                </div>
                <div className="workspace-task">
                  <span className="workspace-task__check workspace-task__check--empty" />
                  <span>Подготовить план запуска</span>
                  <small className="workspace-task__tag">В работе</small>
                </div>
              </div>
            </div>

            <div className="story-quote">
              <div className="story-quote__mark" aria-hidden="true">“</div>
              <blockquote>
                kakoi zhe tyt sverxy zaeb chatgpt spasibo i privet kto eto chitaet
              </blockquote>
              <div className="story-quote__person">
                <div className="story-quote__avatar">КО</div>
                <div>
                  <strong>Отзыв команды</strong>
                  <span>Демонстрационный пример отзыва</span>
                </div>
                <div className="story-quote__rating" aria-label="пять звёзд из пяти">★★★★★</div>
              </div>
            </div>

            <div className="story-panel__bottom">
              <span>Ясность помогает двигаться вперёд.</span>
              <span>01 — 03</span>
            </div>
          </div>
        </aside>
      </div>
    </main>
  );
}
