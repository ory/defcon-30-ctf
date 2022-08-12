import { RecipeVariants } from '@vanilla-extract/recipes';
export declare const buttonStyle: import("@vanilla-extract/recipes/dist/declarations/src/types").RuntimeFn<{
    size: {
        medium: {
            fontSize: string;
            lineHeight: string;
        };
        small: {
            fontSize: string;
            lineHeight: string;
        };
        large: {
            fontSize: string;
            lineHeight: string;
            padding: string;
        };
    };
    variant: {
        regular: {
            fontWeight: number;
            fontStyle: "normal";
        };
        semibold: {
            fontWeight: number;
            fontStyle: "normal";
        };
    };
}>;
export declare type ButtonStyle = RecipeVariants<typeof buttonStyle>;
//# sourceMappingURL=button.css.d.ts.map