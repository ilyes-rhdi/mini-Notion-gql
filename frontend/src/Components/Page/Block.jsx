import { useMemo, useState } from "react";

function getText(data) {
  if (!data) return "";
  if (typeof data === "string") return data;
  if (typeof data === "object" && data.text != null) return String(data.text);
  try {
    return JSON.stringify(data);
  } catch {
    return "";
  }
}

export default function Block({ block, onUpdate, onDelete }) {
  const initial = useMemo(() => getText(block.data), [block.data]);
  const [value, setValue] = useState(initial);

  const typeLabel = (block.Type || "PARAGRAPH").toLowerCase();

  return (
    <div className="ui-card pad-4">
      <div className="row-between">
        <div className="ui-tag">{typeLabel}</div>
        <button className="ui-btn-danger" type="button" onClick={() => onDelete(block.id)}>
          delete
        </button>
      </div>

      <textarea
        className="ui-textarea spacer"
        value={value}
        onChange={(e) => setValue(e.target.value)}
        onBlur={() => onUpdate(block.id, block.Type, { text: value })}
        rows={2}
      />

      <div className="small muted2 spacer">block id: {block.id}</div>
    </div>
  );
}
