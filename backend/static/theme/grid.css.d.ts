import { RecipeVariants } from '@vanilla-extract/recipes';
export declare const gridStyle: import("@vanilla-extract/recipes/dist/declarations/src/types").RuntimeFn<{
    direction: {
        row: {
            flexDirection: "row";
        };
        column: {
            flexDirection: "column";
        };
    };
    gap: {
        4: {
            gap: string;
        };
        8: {
            gap: string;
        };
        16: {
            gap: string;
        };
        32: {
            gap: string;
        };
        64: {
            gap: string;
        };
    };
}>;
export declare type GridStyle = RecipeVariants<typeof gridStyle>;
//# sourceMappingURL=grid.css.d.ts.map