import LoginButton from "./header/LoginButton";
import RegisterButton from "./header/RegisterButton";
import SettingsButton from "./header/SettingsButton";
import LogoutButton from "./header/LogoutButton";
import Logo from "./header/Logo";

interface IHeaderProps {
  authenticated: boolean;
  setAuthenticated: (authenticated: boolean) => void;
}

export default function Header(props: IHeaderProps) {
  return (
    <div className="bg-gray-200 p-4 text-black md:grid md:grid-cols-3">
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
