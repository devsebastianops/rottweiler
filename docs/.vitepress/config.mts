import { defineConfig } from "vitepress";
import { withMermaid } from 'vitepress-plugin-mermaid'

// https://vitepress.dev/reference/site-config
const config = defineConfig({
  title: "Rottweiler",
  description: "Rottweiler is a declarative policy engine that validates any YAML or JSON document against policies you define.",
  
  lang: "en-US",
  cleanUrls: true,
  lastUpdated: true,

  head: [
    ["link", { rel: "icon", sizes: "32x32", href: "/favicon-32x32.png" }],
    ["link", { rel: "icon", sizes: "16x16", href: "/favicon-16x16.png" }],
    ["link", { rel: "apple-touch-icon", sizes: "180x180", href: "/apple-touch-icon.png" }],
    ["link", { rel: "manifest", href: "/site.webmanifest" }],
  ],

  themeConfig: {
    logo: "/assets/rottweiler_mascott.png",

    search: {
      provider: "local"
    },

    // https://vitepress.dev/reference/default-theme-config
    nav: [
       {
        text: "Guide",
        link: "/guide/getting-started"
      },
      {
        text: "Configuration",
        link: "/configuration/overview"
      },
      {
        text: "Examples",
        link: "/examples/overview"
      },
      {
        text: "Reference",
        link: "/reference/cli"
      }
    ],

    sidebar: {
      "/guide/": [
        {
          text: "Guide",
          items: [
            {
              text: "Getting Started",
              link: "/guide/getting-started"
            },
            {
              text: "Installation",
              link: "/guide/installation"
            },
            {
              text: "Quick Start",
              link: "/guide/quick-start"
            },
            {
              text: "Core Concepts",
              link: "/guide/concepts"
            }
          ]
        }
      ],

      "/configuration/": [
        {
          text: "Configuration",
          items: [
            {
              text: "Overview",
              link: "/configuration/overview"
            },
            {
              text: "Rules",
              link: "/configuration/rules"
            },
            {
              text: "Reporting",
              link: "/configuration/reporting"
            },
            {
              text: "Severities",
              link: "/configuration/severity"
            }
          ]
        }
      ],

     

      "/examples/": [
        {
          text: "Examples",
          items: [
            {
              text: "Overview",
              link: "/examples/overview"
            },
            {
              text: "GitHub Workflow Validation",
              link: "/examples/ci-cd-pipeline-validation"
            },
            {
              text: "Docker Compose Check",
              link: "/examples/docker-compose-check"
            },
            {
              text: "Package Governance",
              link: "/examples/package-governance"
            },
            {
              text: "Swagger / OpenAPI Governance",
              link: "/examples/swagger-openapi-governance"
            },
            {
              text: "Terraform Plan Validation",
              link: "/examples/terraform-plan-validation"
            }
          ]
        }
      ],

      "/reference/": [
        {
          text: "Reference",
          items: [
            {
              text: "CLI",
              link: "/reference/cli"
            },
            {
              text: "Configuration",
              link: "/reference/configuration"
            },
            {
              text: "CEL Functions",
              link: "/reference/cel"
            },
            {
              text: "GitHub Action",
              link: "/reference/github-action"
            }
          ]
        }
      ]
    },

    socialLinks: [
      { icon: 'github', link: 'https://github.com/devsebastianops/rottweiler' }
    ],

    footer: {
      message: "Released under the MIT License.",
      copyright: "Copyright © 2026 <a target='_blank' href='https://devsebastianops.com'>Sebastian Breuer</a>"
    },

    editLink: {
      pattern:
        "https://github.com/devsebastianops/rottweiler/edit/main/docs/:path",
      text: "Edit this page on GitHub"
    },

    outline: {
      level: [2, 3],
      label: "On this page"
    },

    docFooter: {
      prev: "Previous",
      next: "Next"
    }
  }
});

export default withMermaid(config);