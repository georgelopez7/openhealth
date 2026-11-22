import type { ErrorComponentProps } from "@tanstack/react-router";

const DefaultErrorLayout = ({ error }: ErrorComponentProps) => {
  const message =
    error instanceof Error ? error.message : "Something went wrong.";

  return (
    <div className="flex min-h-screen flex-col items-center justify-center bg-background p-8 text-center text-white">
      <h1 className="text-4xl font-bold">Error</h1>
      <p className="mt-4 max-w-md text-white/70">{message}</p>
      <button
        type="button"
        onClick={() => window.location.reload()}
        className="mt-6 cursor-pointer text-sm underline underline-offset-4 hover:text-white/80"
      >
        Try again
      </button>
    </div>
  );
};

export default DefaultErrorLayout;
