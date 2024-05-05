interface IOrderPhotosProps {}

export default function OrderPhotos(props: IOrderPhotosProps) {
  return (
    <div>
      <p className="font-medium after:ml-0.5 after:text-red-500 after:content-['*']">
        Photos
      </p>

      <div className="mt-1 flex items-center space-x-6">
        <label className="block">
          <span className="sr-only">Choose profile photo</span>
          <input
            type="file"
            className="block w-full cursor-pointer  text-xs text-gray-400 file:mr-4 file:rounded-full file:border-0 file:bg-violet-50 file:px-4 file:py-2 file:text-sm file:font-semibold file:text-indigo-500 hover:file:bg-indigo-100"
          />
        </label>
      </div>
    </div>
  );
}
