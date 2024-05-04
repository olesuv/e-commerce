import SearchBar from "./SearchBar";

interface IOrganizerProps {
  authenticated: boolean;
}

export default function Organizer(props: IOrganizerProps) {
  return (
    <div className="md:grid md:grid-cols-3 p-4 text-black">
      <div className="hidden md:block"></div>

      <>{props.authenticated ? <SearchBar /> : <p>unauthenticated</p>}</>

      <div className="hidden md:block"></div>
    </div>
  );
}
