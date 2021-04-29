import React from "react";
import { GlobalStyles } from "twin.macro";

const Layout = ({ children, ...rest }) => (
  <div {...rest}>
    <GlobalStyles />
    <div tw="container mx-auto px-4">{children}</div>
    <footer tw="bg-gray-100 mt-8">
      <div tw="container mx-auto px-4 py-8">
        Copyright &copy; Luke Harris {new Date().getFullYear()}
      </div>
    </footer>
  </div>
);

export default Layout;
