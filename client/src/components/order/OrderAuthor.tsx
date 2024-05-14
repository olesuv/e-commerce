import { IAuthorDetailsProps } from "../../../types/orderDetailsProps";

export default function OrderAuthor(props: IAuthorDetailsProps) {
  return (
    <div className="mt-2">
      <p className="text-ms font-medium">Author</p>
      <p className="text-sm text-neutral-500">
        <span className="text-indigo-500 underline">
          {props.authorData?.name}
        </span>{" "}
        with contact{" "}
        <a
          href={`mailto:${props.authorData?.email}`}
          className="text-indigo-500 underline"
        >
          {props.authorData?.email}
        </a>
      </p>
    </div>
  );
}
