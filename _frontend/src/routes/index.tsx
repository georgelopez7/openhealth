import OpenhealthIcon from "#/components/(icons)/openhealth-icon";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/")({ component: Home });

function Home() {
  return (
    <div className="flex min-h-screen items-center justify-center gap-8">
      <OpenhealthIcon className="size-32" />
      <h1 className="text-7xl font-bold">OpenHealth</h1>
    </div>
  );
}
