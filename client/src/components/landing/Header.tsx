import LoginButton from "./LoginButton";
import RegisterButton from "./RegisterButton";
import SettingsButton from "./SettingsButton";
import LogoutButton from "./LogoutButton";
import Logo from "./Logo";

interface IHeaderProps {
  authenticated: boolean;
  setAuthenticated: (authenticated: boolean) => void;
}

export default function Header(props: IHeaderProps) {
  return (
    <div className="md:grid md:grid-cols-3 p-4 bg-gray-200 text-black">
      <div className="hidden md:block"></div>
      <div className="flex justify-between md:items-center">
        <Logo />
        <div>
          {props.authenticated ? (
            <>
              {/*FIX: medium screen responsive*/}
              <SettingsButton />
              <LogoutButton setAuthenticated={props.setAuthenticated} />
            </>
          ) : (
            <>
              {/*FIX: medium screen responsive*/}
              <RegisterButton setAuthenticated={props.setAuthenticated} />
              <LoginButton setAuthenticated={props.setAuthenticated} />
            </>
          )}
        </div>
      </div>
      <div className="hidden md:block"></div>
    </div>
  );
}
