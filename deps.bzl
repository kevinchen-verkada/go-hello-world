load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "com_github_bazelbuild_rules_go",
        importpath = "github.com/bazelbuild/rules_go",
        sum = "h1:ViPR65vOrg74JKntAUFY6qZkheBKGB6to7wFd8gCRU4=",
        version = "v0.35.0",
    )
    go_repository(
        name = "com_github_delta456_box_cli_maker_v2",
        importpath = "github.com/Delta456/box-cli-maker/v2",
        sum = "h1:rGdoK/Qt3shdT1uqRMGgPqrhtisGD7PamTW8vY5MyCA=",
        version = "v2.3.0",
    )
    go_repository(
        name = "com_github_gookit_color",
        importpath = "github.com/gookit/color",
        sum = "h1:uLnfXcaFjlrDnQDT+NCBcfhrXqYTx/rcCa6xn01Y8yI=",
        version = "v1.5.2",
    )
    go_repository(
        name = "com_github_xo_terminfo",
        importpath = "github.com/xo/terminfo",
        sum = "h1:JVG44RsyaB9T2KIHavMF/ppJZNG9ZpyihvCd0w101no=",
        version = "v0.0.0-20220910002029-abceb7e1c41e",
    )
    go_repository(
        name = "org_golang_x_term",
        importpath = "golang.org/x/term",
        sum = "h1:DR4lr0TjUs3epypdhTOkMmuF5CDFJ/8pOnbzMZPQ7bg=",
        version = "v0.32.0",
    )
    go_repository(
        name = "com_github_muesli_reflow",
        importpath = "github.com/muesli/reflow",
        sum = "h1:IFsN6K9NfGtjeggFP+68I4chLZV2yIKsXJFNZ+eWh6s=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_huandu_xstrings",
        importpath = "github.com/huandu/xstrings",
        sum = "h1:2ag3IFq9ZDANvthTwTiqSSZLjDc+BedvHPAp5tJy2TI=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_mattn_go_runewidth",
        importpath = "github.com/mattn/go-runewidth",
        sum = "h1:E5ScNMtiwvlvB5paMFdw9p4kSQzbXFikJ5SQO6TULQc=",
        version = "v0.0.16",
    )
    go_repository(
        name = "com_github_rivo_uniseg",
        importpath = "github.com/rivo/uniseg",
        sum = "h1:WUdvkW8uEhrYfLC4ZzdpI2ztxP1I582+49Oc5Mq64VQ=",
        version = "v0.4.7",
    )
