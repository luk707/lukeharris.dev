import React from "react";
import { GlobalStyles } from "twin.macro";

const Layout = ({ children, ...rest }) => (
  <div {...rest}>
    <GlobalStyles />
    <div tw="container mx-auto px-4">{children}</div>
  </div>
);

export default Layout;
