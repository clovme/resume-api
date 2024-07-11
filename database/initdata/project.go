package initdata

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"resume/models"
	"time"
)

// ProjectExperience 项目经历
func (d *InitData) ProjectExperience() {
	modelList := []models.Project{
		{RID: 0, UID: 0, StartAt: "2017-06", EndAt: "2019-08", Name: "月球开发挖掘机项目", Title: "项目总负责人", Content: "<ol><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>编程语言精通</strong>：熟练掌握多种编程语言，包括但不限于 Go、Python、JavaScript 和 C++。</li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>Web开发</strong>：具备丰富的前端和后端开发经验，精通 Vue.js、React、Node.js 和 Express。</li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>数据库管理</strong>：熟悉 MySQL、PostgreSQL、SQLite 以及 NoSQL 数据库如 MongoDB 的设计、优化与维护。</li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>系统架构设计</strong>：具备设计和实现高可用性、高性能系统架构的能力，了解微服务架构和分布式系统。</li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>项目管理</strong>：具有敏捷开发、Scrum 和 Kanban 等项目管理经验，善于协调团队工作。</li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>版本控制</strong>：精通 Git 及其常用工作流，如 Git Flow，具有团队协作经验。</li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>测试驱动开发</strong>：擅长单元测试、集成测试和功能测试，熟练使用 Jest、JUnit、Go testing 等测试框架。</li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>云计算与部署</strong>：熟悉 AWS、Google Cloud、Azure 等云服务平台，掌握 Docker 和 Kubernetes 等容器化技术。</li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>网络安全</strong>：了解基本的网络安全知识，能够识别和防范常见的安全威胁，如 SQL 注入、XSS 攻击等。</li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>机器学习与数据分析</strong>：具备基本的机器学习知识，熟悉常用算法，能够使用 Pandas、NumPy、TensorFlow 等工具进行数据处理与分析。</li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>代码优化</strong>：擅长代码重构与性能优化，能够有效提升系统的运行效率和可维护性。</li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>文档撰写</strong>：具有良好的技术文档撰写能力，能够清晰地记录项目流程、代码注释和用户手册。</li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>团队合作</strong>：具备优秀的团队合作精神和沟通能力，能够在多文化、多背景的团队中高效工作。</li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>学习能力</strong>：对新技术和新工具保持强烈的好奇心，具备快速学习和应用的能力。</li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>演讲与培训</strong>：拥有良好的演讲能力，能够为团队或客户进行技术培训和知识分享。</li></ol>", ToNow: false, Sort: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	for _, model := range modelList {
		if err := d.Db.Where("rid = ? and uid = ? and name = ?", model.RID, model.UID, model.Name).First(&models.Project{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 数据不存在，插入新记录
				if err := d.Db.Create(&model).Error; err != nil {
					log.Fatalln("[初始化]项目经历创建失败:", err)
				} else {
					log.Println("[初始化]项目经历创建创建成功:", model)
				}
			} else {
				log.Fatalln("[初始化]项目经历查询失败:", err)
			}
		}
	}
}
