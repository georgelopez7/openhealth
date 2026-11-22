import { createFileRoute, Link } from "@tanstack/react-router";
import GitHubIcon from "#/components/(icons)/github-icon";
import OpenhealthIcon from "#/components/(icons)/openhealth-icon";
import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";
import Spacer from "#/components/(layouts)/spacer/spacer";
import { EXTERNAL_LINKS } from "#/domain/config";
import { cn } from "#/lib/utils";

export const Route = createFileRoute("/")({ component: Home });

function Home() {
  return (
    <PageLayout navbar={false}>
      <div className="flex flex-1 flex-col items-center justify-center">
        <div className="flex flex-col items-center gap-3 sm:flex-row sm:gap-6">
          <OpenhealthIcon className="hidden sm:block size-24 md:size-30" />
          <h1 className="text-center text-3xl font-bold sm:text-5xl md:text-6xl">
            OpenHealth
          </h1>
        </div>
        <Spacer size="xsmall" className="my-4 sm:my-1" />
        <div className="flex gap-2">
          <Link
            to="/dashboard"
            className={cn(
              "inline-flex h-11 items-center justify-center gap-3 rounded-lg border border-white/20 bg-white/5 px-4 text-base font-medium text-white transition-colors hover:bg-white/10",
            )}
          >
            Get Started
          </Link>
          <a
            href={EXTERNAL_LINKS["openhealth.github"]}
            target="_blank"
            rel="noopener noreferrer"
            className={cn(
              "inline-flex size-11 items-center justify-center rounded-lg border border-white/20 bg-white/5 text-white transition-colors hover:bg-white/10",
            )}
            aria-label="GitHub"
          >
            <GitHubIcon className="size-5" />
          </a>
        </div>
        <Spacer size="medium" />
        <div className="inline-flex items-center gap-2 px-4 py-2 text-[10px] text-white/60">
          <span>
            Powered by{" "}
            <a
              href={EXTERNAL_LINKS["openfga.website"]}
              target="_blank"
              rel="noopener noreferrer"
              className="font-semibold underline sm:no-underline sm:hover:underline underline-offset-2 transition-colors hover:text-white"
            >
              OpenFGA
            </a>
          </span>
        </div>
        <Spacer size="medium" />
        <OpenhealthIcon className="sm:hidden size-32 opacity-70" />
      </div>
    </PageLayout>
  );
}
