import { QueryClient } from "@tanstack/react-query";
import { createRouter as createTanStackRouter } from "@tanstack/react-router";
import { setupRouterSsrQueryIntegration } from "@tanstack/react-router-ssr-query";
import DefaultErrorLayout from "./components/(layouts)/default-error-layout/default-error-layout";
import NotFoundLayout from "./components/(layouts)/not-found-layout/not-found-layout";
import { routeTree } from "./routeTree.gen";

export function getRouter() {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        refetchOnWindowFocus: false,
        staleTime: 1000 * 60,
      },
    },
  });

  const router = createTanStackRouter({
    routeTree,
    context: { queryClient },
    defaultErrorComponent: DefaultErrorLayout,
    defaultNotFoundComponent: NotFoundLayout,
    notFoundMode: "root",
  });

  setupRouterSsrQueryIntegration({ router, queryClient });

  return router;
}
