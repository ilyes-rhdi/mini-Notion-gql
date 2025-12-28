import { useEffect, useMemo, useState } from "react";

import Block from "./Block";

export default function Editor({
  page,
  blocks,
  onRenamePage,
  onCreateParagraph,
  onUpdateBlock,
  onDeleteBlock,
}) {
  const safeBlocks = useMemo(() => blocks || [], [blocks]);

  const [title, setTitle] = useState(() => page?.title ?? "");

  useEffect(() => {
    setTitle(page?.title ?? "");
  }, [page?.id]);

  if (!page) {
    return (
      <div className="ui-panel pad-6">
        <div className="muted">select a page on the left</div>
      </div>
    );
  }

  return (
    <div className="stack">
      <div className="ui-panel pad-5">
        <div className="row-between">
          <div className="ui-tag">page</div>
          <button className="ui-btn-secondary" type="button" onClick={onCreateParagraph}>
            + paragraph
          </button>
        </div>

        <input
          className="title-input"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          onBlur={() => onRenamePage(page.id, title)}
          placeholder="Untitled"
        />

        <div className="small muted2">page id: {page.id}</div>
      </div>

      <div className="stack-sm">
        {safeBlocks.map((b) => (
          <Block key={b.id} block={b} onUpdate={onUpdateBlock} onDelete={onDeleteBlock} />
        ))}

        {!safeBlocks.length ? (
          <div className="ui-card pad-4 muted">
            no blocks yet
          </div>
        ) : null}
      </div>
    </div>
  );
}
