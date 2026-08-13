import { createFileRoute, Link } from "@tanstack/react-router";
import GitHubIcon from "#/components/(icons)/github-icon";
import OpenhealthIcon from "#/components/(icons)/openhealth-icon";
import Spacer from "#/components/(layouts)/spacer/spacer";
import { buttonVariants } from "#/components/ui/button";
import { cn } from "#/lib/utils";

export const Route = createFileRoute("/")({ component: Home });

const buttonBaseClasses = "h-12 border-2 border-dashed border-white text-lg";

function Home() {
  return (
    <div className="flex min-h-screen flex-col items-center justify-center">
      <div className="flex items-center gap-6">
        <OpenhealthIcon className="size-30" />
        <h1 className="text-6xl font-bold">OpenHealth</h1>
      </div>
      <Spacer size="xsmall" />
      <div className="flex gap-3">
        <Link
          to="/dashboard"
          className={cn(
            buttonVariants({ variant: "outline", size: "lg" }),
            buttonBaseClasses,
            "px-6",
          )}
        >
          Get Started
        </Link>
        <a
          href="https://github.com"
          target="_blank"
          rel="noopener noreferrer"
          className={cn(
            buttonVariants({ variant: "outline", size: "icon-lg" }),
            buttonBaseClasses,
            "size-12",
          )}
        >
          <GitHubIcon className="size-6" />
        </a>
      </div>
    </div>
  );
}
