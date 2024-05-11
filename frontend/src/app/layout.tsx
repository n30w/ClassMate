import type { Metadata } from "next";
import "./globals.css";
import localFont from "next/font/local";

const neue = localFont({
  src: [
    {
      path: "/fonts/NeueMontreal-Regular.woff2",
      style: "regular",
    },
    {
      path: "/fonts/NeueMontreal-Bold.woff2",
      style: "bold",
    },
    {
      path: "/fonts/NeueMontreal-Light.woff2",
      style: "light",
    },
    {
      path: "/fonts/NeueMontreal-Medium.woff2",
      style: "medium",
    },
  ],
});

export const metadata: Metadata = {
  title: "Darkspace",
  description: "NYU Darkspace",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${neue.className}`}>{children}</body>
    </html>
  );
}
