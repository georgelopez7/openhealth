import { Link } from "@tanstack/react-router";

interface INotFoundLayoutProps {
  title?: string;
  message?: string;
  data?: unknown;
  isNotFound?: boolean;
  routeId?: string;
}

const NotFoundLayout = ({
  title = "404",
  message = "Page not found.",
}: INotFoundLayoutProps) => {
  return (
    <div className="flex min-h-screen flex-col items-center justify-center bg-background p-8 text-center text-white">
      <h1 className="text-4xl font-bold">{title}</h1>
      <p className="mt-4 text-white/70">{message}</p>
      <Link
        to="/"
        className="mt-6 text-sm underline underline-offset-4 hover:text-white/80"
      >
        Go back home
      </Link>
    </div>
  );
};

export default NotFoundLayout;
