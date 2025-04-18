import {
  createBrowserRouter,
  RouterProvider,
  redirect,
} from "react-router-dom";
import HomePage from "../pages/index";
import LoginPage from "../pages/auth/login";
import BaseSimplePage from "../layouts/base";
import RegisterPage from "@/pages/auth/register";
import RequireAuth from "./RequireAuth";
import BaseLayout from "@/layouts/backstage";
import SelectWorkspace from "@/pages/workspace/select";
import WorkspaceNote from "@/pages/workspace/note";
let router = createBrowserRouter([
  {
    path: "/",
    element: (
      <RequireAuth>
        <BaseLayout />
      </RequireAuth>
    ),
    children: [
      {
        path: "",
        index: true,
        Component: LoginPage,
      },
      {
        path: "workspace/:id",
        index: true,
        Component: WorkspaceNote,
      },
    ],
  },
  {
    path: "/",
    element: (
      <RequireAuth>
        <BaseSimplePage />
      </RequireAuth>
    ),
    children: [
      {
        path: "select",
        index: true,
        Component: SelectWorkspace,
      },
    ],
  },
  {
    path: "/auth",
    Component: BaseSimplePage,
    children: [
      {
        path: "login",
        index: true,
        Component: LoginPage,
      },
      {
        path: "register",
        Component: RegisterPage,
      },
    ],
  },
]);

function AppRouter() {
  return <RouterProvider router={router}></RouterProvider>;
}

export default AppRouter;
