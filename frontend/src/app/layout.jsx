import { Geist, Geist_Mono, Krona_One, Space_Mono} from "next/font/google";
import "./globals.css";

import Header from "@/components/ui/Header/Header";

const spaceMono = Space_Mono({
  weight: '400',
  variable: '--font-space-mono',
  subsets: ['latin'],
})

const kronaOne = Krona_One({
  weight: '400',
  variable: '--font-krona-one',
  subsets: ["latin"]
})

const geistSans = Geist({
  weight: '400',
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

export const metadata = {
  title: "Efficax Counter",
  description: "A simple counter of numbers. - Create by dvcDaniel.",
};

export default function RootLayout({ children }) {
  return (
    <html
      lang="en"
      className={`${geistSans.variable} ${kronaOne.variable} ${spaceMono.variable} h-full antialiased`}
    >
      <body>
         <Header/>
          {children}
      </body>
    </html>
  );
}
